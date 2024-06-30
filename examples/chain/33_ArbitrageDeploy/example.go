package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
	strategytypes "github.com/FluxNFTLabs/sdk-go/chain/modules/strategy/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "embed"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	//go:embed arbitrage_demo.wasm
	botBinary []byte
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
	cc, err := grpc.Dial("localhost:9900", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	msg := &strategytypes.MsgConfigStrategy{
		Sender:   senderAddress.String(),
		Config:   strategytypes.Config_deploy,
		Strategy: botBinary,
		Query:    &types.FISQueryRequest{},
		Metadata: &strategytypes.StrategyMetadata{
			Name:        "Amm swap wizard",
			Description: "Abstracting out swaps and unleashing arbitrage opportunities. It supports basic swaps as well as arbitrage across pools of the same pair",
			Logo:        "https://media.licdn.com/dms/image/D560BAQG83Lxt5OyW4Q/company-logo_200_200/0/1706516333778?e=1724889600&v=beta&t=_fD4VtBEpNKJ8gESgFOcZyTDrQFlRUfxw1iQD8mItnM",
			Website:     "https://www.astromesh.xyz",
			Type:        strategytypes.StrategyType_INTENT_SOLVER,
			Tags:        []string{"helper", "defi"},
			Schema:      "",
		},
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	res, err := chainClient.SyncBroadcastMsg(msg)
	if err != nil {
		panic(err)
	}

	hexResp, err := hex.DecodeString(res.TxResponse.Data)
	if err != nil {
		panic(err)
	}

	// decode result to get contract address
	var txData sdk.TxMsgData
	if err := txData.Unmarshal(hexResp); err != nil {
		panic(err)
	}

	var r strategytypes.MsgConfigStrategyResponse
	if err := r.Unmarshal(txData.MsgResponses[0].Value); err != nil {
		panic(err)
	}

	fmt.Println("gas used/want:", res.TxResponse.GasUsed, "/", res.TxResponse.GasWanted)
	fmt.Println("strategy id:", r.Id)
}
