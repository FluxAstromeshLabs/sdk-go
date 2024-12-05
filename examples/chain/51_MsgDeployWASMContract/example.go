package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "embed"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
)

var (
	//go:embed wasm_contract.wasm
	wasmBytes []byte
)

func main() {
	valAddr, err := sdk.ValAddressFromBech32("luxvaloper1qry5x2d383v9hkqc0fpez53yluyxvey2c957m4")
	if err != nil {
		panic(err)
	}

	consAddrPub, err := base64.StdEncoding.DecodeString("uQi1wxB44tBRDWRU4h1IhdWy7RIB2gXe/v8jdthNbeE=")
	if err != nil {
		panic(err)
	}

	fmt.Println("val addr:", hex.EncodeToString(valAddr), hex.EncodeToString(tmhash.Sum(consAddrPub)))

	return
	network := common.LoadNetwork("local", "")
	kr, err := keyring.New(
		"fluxd",
		"file",
		os.Getenv("HOME")+"/.fluxd",
		strings.NewReader("12345678\n"),
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
		panic(err)
	}

	storeCodeMsg := &wasmtypes.MsgStoreCode{
		Sender:       senderAddress.String(),
		WASMByteCode: wasmBytes,
	}

	res, err := chainClient.SyncBroadcastMsg(storeCodeMsg)
	if err != nil {
		fmt.Println(err)
	}

	var codeID uint64

	for _, e := range res.TxResponse.Events {
		if e.Type == "store_code" {
			for _, attr := range e.Attributes {
				if string(attr.Key) == "code_id" {
					codeID, err = strconv.ParseUint(string(attr.Value), 10, 64)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}

	instantiateMsg := &wasmtypes.MsgInstantiateContract{
		Sender: senderAddress.String(),
		CodeID: codeID,
		Label:  "wasm_contract",
		Msg:    []byte("{}"),
		Funds:  nil,
	}

	txResp, err := chainClient.SyncBroadcastMsg(instantiateMsg)
	if err != nil {
		panic(err)
	}

	fmt.Println("tx hash:", txResp.TxResponse.TxHash)
	fmt.Println("gas used/want:", txResp.TxResponse.GasUsed, "/", txResp.TxResponse.GasWanted)
}
