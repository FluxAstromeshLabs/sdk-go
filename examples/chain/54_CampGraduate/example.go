package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ethsecp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
)

type Project struct {
	Denom      string `json:"denom"`
	Graduated  bool   `json:"graduated"`
	PoolSeeded bool   `json:"pool_seeded"`
	Creator    string `json:"creator"`
}

func graduateProject(chainClient chainclient.ChainClient, projects []*Project) error {
	for _, project := range projects {
		if project.Graduated == true && project.PoolSeeded == false {
			graduateMsg := map[string]interface{}{
				"graduate": map[string]interface{}{
					"camp_denom": project.Denom,
				},
			}

			graduateBz, _ := json.Marshal(graduateMsg)
			msg := &wasmtypes.MsgExecuteContract{
				Sender:   project.Creator,
				Contract: "inj1wkwy0xh89ksdgj9hr347dyd2dw7zesmtcf6an0",
				Msg:      graduateBz,
			}

			res, err := chainClient.SyncBroadcastMsg(msg)
			if err != nil {
				return fmt.Errorf("failed to graduate project: %w", err)
			}

			fmt.Println("graduate ", project.Denom)
			fmt.Println("tx hash", res.TxResponse.TxHash)
			fmt.Println("gas used/want:", res.TxResponse.GasUsed, "/", res.TxResponse.GasWanted)
		}
	}
	return nil
}

func injectiveConfig() {
	cfg := sdk.GetConfig()
	chaintypes.SetBech32Prefixes(cfg)
	cfg.SetBech32PrefixForAccount("inj", "injpub")
	proto.RegisterType((*ethsecp256k1.PubKey)(nil), "injective.crypto.v1beta1.ethsecp256k1.PubKey")
	proto.RegisterType((*chaintypes.EthAccount)(nil), "injective.types.v1beta1.EthAccount")
}

func main() {
	network := common.LoadNetwork("local", "")
	network.ChainId = "injective-888"

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

	injectiveConfig()

	// init grpc connection
	cc, err := grpc.Dial(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// init client ctx
	clientCtx, _, err := chaintypes.NewClientContext(
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
		common.OptionGasPrices("500000000inj"),
	)
	if err != nil {
		fmt.Println(err)
	}

	//smart state query
	wasmClient := wasmtypes.NewQueryClient(cc)

	for {
		queryMsg := map[string]interface{}{
			"projects": map[string]interface{}{},
		}

		queryBz, _ := json.Marshal(queryMsg)

		projectsBz, err := wasmClient.SmartContractState(context.Background(), &wasmtypes.QuerySmartContractStateRequest{
			Address:   "inj1wkwy0xh89ksdgj9hr347dyd2dw7zesmtcf6an0",
			QueryData: queryBz,
		})
		if err != nil {
			panic(err)
		}

		var projects []*Project
		err = json.Unmarshal(projectsBz.Data, &projects)
		if err != nil {
			panic(err)
		}

		err = graduateProject(chainClient, projects)
		if err != nil {
			panic(err)
		}

		time.Sleep(5 * time.Second)
	}
}
