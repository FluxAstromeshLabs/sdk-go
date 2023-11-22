package main

import (
	sdkmath "cosmossdk.io/math"
	"fmt"
	bazaartypes "github.com/FluxNFTLabs/sdk-go/chain/modules/bazaar/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"os"
	"strings"

	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("local", "")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

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

	// initialize grpc client
	clientCtx, senderAddress, err := chaintypes.NewClientContext(
		network.ChainId,
		"genesis",
		kr,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	// prepare tx msg
	content := &bazaartypes.ClassCommissionsProposal{
		Title:       "Update class commission proposal",
		Description: "09.15.2023",
		ClassCommissions: []bazaartypes.ClassCommission{
			{ClassId: "series", CommissionMul: 77, CommissionDiv: 100},
			{ClassId: "livestream", CommissionMul: 45, CommissionDiv: 100},
		},
	}
	contentAny, err := codectypes.NewAnyWithValue(content)
	if err != nil {
		panic(err)
	}

	depositAmount, _ := sdkmath.NewIntFromString("500000000000000000000")
	proposalMsg := &govtypes.MsgSubmitProposal{
		Content:        contentAny,
		InitialDeposit: sdk.Coins{{Denom: "lux", Amount: depositAmount}},
		Proposer:       senderAddress.String(),
	}

	// prepare vote msg using node account
	voteMsg := &govtypes.MsgVote{
		ProposalId: 1,
		Voter:      senderAddress.String(),
		Option:     govtypes.OptionYes,
	}

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000lux"),
	)
	if err != nil {
		fmt.Println(err)
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(proposalMsg)
	if err != nil {
		fmt.Println(err)
	}
	chainClient.BroadcastDone()

	err = chainClient.QueueBroadcastMsg(voteMsg)
	if err != nil {
		fmt.Println(err)
	}
	chainClient.BroadcastDone()
}
