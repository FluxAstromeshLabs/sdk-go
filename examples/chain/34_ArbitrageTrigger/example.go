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
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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

	// prepare tx msg
	if err != nil {
		panic(err)
	}

	senderSvm := solana.PublicKeyFromBytes(ethcrypto.Keccak256(senderAddress.Bytes()))
	msg := &strategytypes.MsgTriggerStrategies{
		Sender: senderAddress.String(),
		Ids:    []string{"f7534ee745ff600d5af41c7e67013f17f78a040ccb48d17b4b39d37238af866f"},
		Inputs: [][]byte{
			[]byte(fmt.Sprintf(`[
				{
					"dex_name": "astroport",
					"pool_id": "lux1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqhywrts",
					"input_denom": "usdt",
					"output_denom": "btc",
					"denom_plane": "COSMOS",
					"input_amount": "10000000"
				}, 
				{
					"dex_name": "raydium",
					"pool_id": "5qUshuBSTpuMu5c1C1Fxq8uJ7Emhn9AAtQwVJfEXAPmy",
					"input_denom": "ENyus6yS21v95sreLKcVEA5Wjcyh8jg6w4jBFHzJaPox",
					"output_denom": "ErDYXZUZ9rpSSvdWvrsQwgh6K4BQeoY2CPyv1FeD1S9r",
					"denom_plane": "SVM",
					"raydium_accounts": {
						"sender_svm_account": "%s",
						"authority_account": "GpMZbSM2GgvTKHJirzeGfMFoaZ8UR2X7F4v8vHTvxFbL",
						"amm_config_account": "D4FPEruKEHrG5TenZ2mpDGEfu1iUvTiqBxvpU8HLBvC2",
						"pool_state_account": "HUtjobntUDzrsq1k7xLM6SLzuZyUvr2U8skA8SUWevFd",
						"input_token_account": "9kWnPUAkspGW6qGPPah1aAdH316nkiJhow5neRs5YDej",
						"output_token_account": "HwkqUQaXocRwNLGX2qKmC3Sk4uTVxmzmCEAEHDwSj4KQ",
						"input_vault": "9U5Lpfmc6u1rCRAfzGe883KnK5Avm76zX4te6sexvCEk",
						"output_vault": "UURmKznoUTh8Dt9wgyusq6u1ETuY8Zj79NFAtfQJ7HB",
						"observer_state": "FXqXrt2xDrxg7J5wdXrTbB2hCGajSzXLvwvc4x3Uw7i"
					}
				}
			]`, senderSvm.String())),
		},
		Queries: []*astromeshtypes.FISQueryRequest{
			{
				Instructions: []*astromeshtypes.FISQueryInstruction{
					{
						Plane:   astromeshtypes.Plane_WASM,
						Action:  astromeshtypes.QueryAction_VM_QUERY,
						Address: sdk.MustAccAddressFromBech32("lux1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqhywrts"),
						Input: [][]byte{
							[]byte(`{"pool":{}}`),
						},
					},
					{
						Plane:  astromeshtypes.Plane_SVM,
						Action: astromeshtypes.QueryAction_VM_QUERY,
						Input: [][]byte{
							solana.MustPublicKeyFromBase58("9U5Lpfmc6u1rCRAfzGe883KnK5Avm76zX4te6sexvCEk").Bytes(), // denom 0
							solana.MustPublicKeyFromBase58("UURmKznoUTh8Dt9wgyusq6u1ETuY8Zj79NFAtfQJ7HB").Bytes(),  // denom 1
						},
					},
				},
			},
		},
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	res, err := chainClient.SyncBroadcastMsg(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
