package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cosmossdk.io/math"
	astromeshtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
	strategytypes "github.com/FluxNFTLabs/sdk-go/chain/modules/strategy/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cosmos/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/gagliardetto/solana-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/FluxNFTLabs/sdk-go/chain/indexer/explorer"
)

func main() {
	network := common.LoadNetwork("local", "")
	kr, err := keyring.New(
		"fluxd",
		"file",
		os.Getenv("HOME")+"/.fluxd",
		strings.NewReader("12345678"),
		chainclient.GetCryptoCodec(),
	)
	if err != nil {
		panic(err)
	}

	// init grpc connection
	cc, err := grpc.Dial(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// init client ctx
	clientCtx, senderAddress, err := chaintypes.NewClientContext(
		network.ChainId,
		"signer4",
		kr,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithGRPCClient(cc)

	// init chain client
	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		common.OptionGasPrices("500000000lux"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("connecting to indexer ...")
	indexerConn, err := grpc.Dial("localhost:4474", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer cc.Close()
	if err != nil {
		panic(err)
	}
	explorerClient := explorer.NewAPIClient(indexerConn)

	fmt.Println("sender address:", senderAddress.String())
	driftProgramId := solana.MustPublicKeyFromBase58("FLR3mfYrMZUnhqEadNJVwjUhjX8ky9vE9qTtDmkK4vwC")
	// check if account is linked, not then create
	isSvmLinked, svmPubkey, err := chainClient.GetSVMAccountLink(context.Background(), senderAddress)
	if err != nil {
		panic(err)
	}

	if !isSvmLinked {
		svmKey := ed25519.GenPrivKey() // Good practice: Backup this private key
		res, err := chainClient.LinkSVMAccount(svmKey, math.NewIntFromUint64(1000_000_000_000))
		if err != nil {
			panic(err)
		}
		fmt.Println("linked sender to svm address:", base58.Encode(svmKey.PubKey().Bytes()), "txHash:", res.TxResponse.TxHash)
	} else {
		fmt.Println("sender is already linked to svm address:", svmPubkey.String())
	}

	fillerUser, _, err := solana.FindProgramAddress([][]byte{
		[]byte("user"), svmPubkey[:], {0, 0},
	}, driftProgramId)
	if err != nil {
		panic(err)
	}

	priceLimit := 64500_000_000
	quantity := 1000000
	fillableOrderResp, err := explorerClient.ListFillableDriftJITOrders(
		context.Background(), &explorer.ListFillableDriftJITOrdersRequest{
			MarketName: "btc-usdt",
			WorstPrice: uint64(priceLimit),
			Direction:  "short",
			Quantity:   uint64(quantity),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Going to fill against", len(fillableOrderResp.FillableOrders), "order(s)")
	if len(fillableOrderResp.FillableOrders) == 0 {
		fmt.Println("Nothing to fill, exit")
		return
	}

	inputs := [][]byte{}
	queries := []*astromeshtypes.FISQueryRequest{}
	strategyId := "19385c1578a1d31b9bded3d83cce11b3fcfb283ac0a5d518b3619c3728d55a7b"

	strategyIds := []string{}
	for _, o := range fillableOrderResp.FillableOrders {
		fillMsg := fmt.Sprintf(
			`{"fill_perp_market_order":{"taker_svm_address":"%s","taker_order_id":"%d","quantity":"%d"}}`,
			o.OwnerAddress, o.OrderId, o.FillableQuantity,
		)
		takerUser, _, err := solana.FindProgramAddress([][]byte{
			[]byte("user"), solana.MustPublicKeyFromBase58(o.OwnerAddress).Bytes(), {0, 0},
		}, driftProgramId)
		if err != nil {
			panic(err)
		}
		strategyIds = append(strategyIds, strategyId)
		inputs = append(inputs, []byte(fillMsg))
		queries = append(queries, &astromeshtypes.FISQueryRequest{
			Instructions: []*astromeshtypes.FISQueryInstruction{
				{
					Plane:   astromeshtypes.Plane_COSMOS,
					Action:  astromeshtypes.QueryAction_COSMOS_QUERY,
					Address: nil,
					Input: [][]byte{
						[]byte("/flux/svm/v1beta1/account_link/cosmos/" + senderAddress.String()),
					},
				},
				{
					Plane:   astromeshtypes.Plane_SVM,
					Action:  astromeshtypes.QueryAction_VM_QUERY,
					Address: nil,
					Input: [][]byte{
						fillerUser[:],
						takerUser[:],
					},
				},
			},
		})
	}

	msgTriggerStategy := &strategytypes.MsgTriggerStrategies{
		Sender:  senderAddress.String(),
		Ids:     strategyIds,
		Inputs:  inputs,
		Queries: queries,
	}
	res, err := chainClient.SyncBroadcastMsg(msgTriggerStategy)
	if err != nil {
		panic(err)
	}

	fmt.Println("tx hash:", res.TxResponse.TxHash)
	fmt.Println("gas used/want:", res.TxResponse.GasUsed, "/", res.TxResponse.GasWanted)
}
