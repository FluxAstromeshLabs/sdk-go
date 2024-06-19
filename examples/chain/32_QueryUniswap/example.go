package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	astromeshtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
	evmtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/evm/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "embed"

	"cosmossdk.io/math"
	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	PoolManager = "07aa076883658b7ed99d25b1e6685808372c8fe2"

	//go:embed PoolManager.json
	PoolManagerCompData []byte
	PoolManagerABI, _   = loadABI(PoolManagerCompData)
	Six                 = ethcommon.LeftPadBytes([]byte{6}, 32)
)

func loadABI(compDataBz []byte) (abi.ABI, error) {
	var compData map[string]interface{}
	err := json.Unmarshal(compDataBz, &compData)
	if err != nil {
		return abi.ABI{}, err
	}
	abiBz, err := json.Marshal(compData["abi"].([]interface{}))
	if err != nil {
		return abi.ABI{}, err
	}
	return abi.JSON(bytes.NewReader(abiBz))
}

type PairKey struct {
	Currency0   ethcommon.Address // base denom addr
	Currency1   ethcommon.Address // quote denom addr
	Fee         *big.Int          // 3000 -> 0.3%
	TickSpacing *big.Int
	Hooks       ethcommon.Address // hook contract addr
}

func (p *PairKey) PoolId() ethcommon.Hash {
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		panic(err)
	}

	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(err)
	}

	tupleAbi := abi.Arguments{
		{
			Type: addressType,
		},
		{
			Type: addressType,
		},
		{
			Type: uint256Type,
		},
		{
			Type: uint256Type,
		},
		{
			Type: addressType,
		},
	}

	cd, err := tupleAbi.Pack(p.Currency0, p.Currency1, p.Fee, p.TickSpacing, p.Hooks)
	if err != nil {
		panic(err)
	}

	return crypto.Keccak256Hash(cd)
}

type PoolInfo struct {
	SqrtPriceX96 math.Int
	Tick         math.Int
	ProtocolFee  uint32
	LpFee        uint32
}

func signedBigIntFromBytes(b []byte) *big.Int {
	res := new(big.Int).SetBytes(b)
	// first bit is set, then it's negative
	// we sub the complement part to get the origin negative number
	if b[0]&0x80 != 0 {
		bitCount := len(b) * 8
		complement := new(big.Int).Lsh(big.NewInt(1), uint(bitCount))
		res = new(big.Int).Sub(res, complement)
	}

	return res
}

// Define the function to get pool info from the input bytes32 data
func parsePoolinfo(data []byte) (*PoolInfo, error) {
	if len(data) != 32 {
		return nil, fmt.Errorf("input data must be 32 bytes")
	}

	// Extract the bottom 160 bits for sqrtPriceX96
	sqrtPriceX96 := new(big.Int).SetBytes(data[12:32])

	// Extract the next 24 bits for tick
	tickBytes := data[9:12]
	tick := signedBigIntFromBytes(tickBytes)

	// Extract the next 24 bits for protocolFee
	protocolFeeBytes := data[6:9]
	protocolFee := new(big.Int).SetBytes(protocolFeeBytes).Uint64()

	// Extract the last 24 bits for lpFee
	lpFeeBytes := data[3:6]
	lpFee := new(big.Int).SetBytes(lpFeeBytes).Uint64()

	return &PoolInfo{
		SqrtPriceX96: math.NewIntFromBigInt(sqrtPriceX96),
		Tick:         math.NewIntFromBigInt(tick),
		ProtocolFee:  uint32(protocolFee),
		LpFee:        uint32(lpFee),
	}, nil
}

func getDenomLink(astromeshClient astromeshtypes.QueryClient, cosmosDenom string) string {
	denomLink, err := astromeshClient.DenomLink(context.Background(), &astromeshtypes.QueryDenomLinkRequest{
		SrcPlane: astromeshtypes.Plane_COSMOS,
		DstPlane: astromeshtypes.Plane_EVM,
		SrcAddr:  cosmosDenom,
	})
	if err != nil {
		panic(err)
	}

	return denomLink.DstAddr
}

func main() {
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

	// // init chain client
	// chainClient, err := chainclient.NewChainClient(
	// 	clientCtx,
	// 	common.OptionGasPrices("500000000lux"),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// read bytecode
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	poolManagerBz, err := os.ReadFile(dir + "/examples/chain/28_DeployUniswap/PoolManager.json")
	if err != nil {
		panic(err)
	}

	var compData map[string]interface{}
	err = json.Unmarshal(poolManagerBz, &compData)
	if err != nil {
		panic(err)
	}
	poolManagerBytecode, err := hex.DecodeString(compData["bytecode"].(map[string]interface{})["object"].(string))
	if err != nil {
		panic(err)
	}

	abiBz, err := json.Marshal(compData["abi"].([]interface{}))
	if err != nil {
		panic(err)
	}

	poolManagerABI, err := abi.JSON(strings.NewReader(string(abiBz)))
	if err != nil {
		panic(err)
	}

	_ = poolManagerBytecode
	_ = poolManagerABI
	_ = senderAddress

	evmQc := evmtypes.NewQueryClient(cc)
	astroQc := astromeshtypes.NewQueryClient(cc)

	btcEvm := getDenomLink(astroQc, "btc")
	usdtEvm := getDenomLink(astroQc, "usdt")
	denom0, denom1 := btcEvm, usdtEvm
	if denom0 > denom1 {
		denom1, denom0 = denom0, denom1
	}

	p := &PairKey{
		Currency0:   ethcommon.HexToAddress(denom0),
		Currency1:   ethcommon.HexToAddress(denom1),
		Fee:         big.NewInt(3000),
		TickSpacing: big.NewInt(60),
		Hooks:       ethcommon.HexToAddress("0x"),
	}

	fmt.Println("poolId:", p.PoolId())

	ctx := context.Background()
	poolId := p.PoolId()
	slot0 := crypto.Keccak256Hash(append([]byte(poolId[:]), Six...))
	fmt.Println("slot 0:", slot0)

	slot0Cd, err := PoolManagerABI.Pack("extsload", slot0)
	if err != nil {
		panic(err)
	}

	res, err := evmQc.ContractQuery(ctx, &evmtypes.ContractQueryRequest{
		Calldata: slot0Cd,
		Address:  PoolManager,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("output:", res.Output)
	poolInfo, err := parsePoolinfo(res.Output)
	if err != nil {
		panic(err)
	}

	bz, _ := json.Marshal(poolInfo)
	fmt.Println("poolInfo:", string(bz))
}
