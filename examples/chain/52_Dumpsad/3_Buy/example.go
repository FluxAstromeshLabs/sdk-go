package main

import (
	"fmt"
	"os"
	"strings"

	astromeshtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
	strategytypes "github.com/FluxNFTLabs/sdk-go/chain/modules/strategy/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
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
		"user2",
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

	// prepare tx msg
	if err != nil {
		panic(err)
	}

	fmt.Println("sender address:", senderAddress.String())
	id := "c666878dd41c8594615060332b1d8171eb754126e690d68857c6620a56510bab"
	poolAddress := "lux1z5mjdzgkqlpqs7rr97tk4xe83ds3wtvnw4advq"
	msgTriggerStategy := &strategytypes.MsgTriggerStrategies{
		Sender: senderAddress.String(),
		Ids:    []string{id},
		Inputs: [][]byte{
			[]byte(
				fmt.Sprintf(
					`{"buy":{"denom":"astromesh/lux1jcltmuhplrdcwp7stlr4hlhlhgd4htqhu86cqx/chill-guy","amount":"2000000000","slippage":"2000","pool_address":"%s"}}`, poolAddress,
				),
			),
		},
		Queries: []*astromeshtypes.FISQueryRequest{
			{
				Instructions: []*astromeshtypes.FISQueryInstruction{
					{
						Plane:   astromeshtypes.Plane_COSMOS,
						Action:  astromeshtypes.QueryAction_COSMOS_BANK_BALANCE,
						Address: nil,
						Input: [][]byte{
							[]byte(poolAddress + "," + poolAddress),
							[]byte("sol,astromesh/lux1jcltmuhplrdcwp7stlr4hlhlhgd4htqhu86cqx/chill-guy"),
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
