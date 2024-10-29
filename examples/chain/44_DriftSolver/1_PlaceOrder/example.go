package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	astromeshtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
	strategytypes "github.com/FluxNFTLabs/sdk-go/chain/modules/strategy/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/gagliardetto/solana-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		"user1",
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
		fmt.Println(err)
	}

	// prepare tx msg
	if err != nil {
		panic(err)
	}

	fmt.Println("sender address:", senderAddress.String())

	driftProgramId := solana.MustPublicKeyFromBase58("FLR3mfYrMZUnhqEadNJVwjUhjX8ky9vE9qTtDmkK4vwC")

	isSvmLinked, SvmPubkey, err := chainClient.GetSVMAccountLink(context.Background(), senderAddress)
	if err != nil {
		panic(err)
	}

	if !isSvmLinked {
		panic(fmt.Errorf("taker is not linked: %s", senderAddress.String()))
	}

	User, _, err := solana.FindProgramAddress([][]byte{
		[]byte("user"), SvmPubkey[:], {0, 0},
	}, driftProgramId)
	if err != nil {
		panic(err)
	}

	marketPubkey0 := solana.MustPublicKeyFromBase58("7WrZxBiKCMGuzLCW2VwKK7sQjhTZLbDe5sKfJsEcARpF")
	marketPubkey1 := solana.MustPublicKeyFromBase58("E4DJDZwcSWzujRLjoWQXqq4KQVuzbvBiHSR35BPbK7BX")
	marketPubkey2 := solana.MustPublicKeyFromBase58("2GKUdmaBJNjfCucDT14HrsWchVrm3yvj4QY2jjnUEg3v")

	msgTriggerStategy := &strategytypes.MsgTriggerStrategies{
		Sender: senderAddress.String(),
		Ids:    []string{"7b1ba16fb3eccb0654881bb25d8e0f7f8370657f5de9815d8a19aff49b542509"},
		Inputs: [][]byte{
			[]byte(`{"place_perp_market_order":{"usdt_amount":"24356000","leverage":5,"market":"eth-usdt","auction_duration":10}}`),
		},
		Queries: []*astromeshtypes.FISQueryRequest{
			{
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
							User[:],
							marketPubkey0[:],
							marketPubkey1[:],
							marketPubkey2[:],
						},
					},
				},
			},
		},
	}

	res, err := chainClient.SyncBroadcastMsg(msgTriggerStategy)
	if err != nil {
		panic(err)
	}

	fmt.Println("tx hash:", res.TxResponse.TxHash)
	fmt.Println("gas used/want:", res.TxResponse.GasUsed, "/", res.TxResponse.GasWanted)
}