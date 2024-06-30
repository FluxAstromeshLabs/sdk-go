package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	gomath "math"
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

	//go:embed erc20.json
	Erc20Bz     []byte
	Erc20ABI, _ = loadABI(Erc20Bz)
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

func getStateSlot(poolId ethcommon.Hash) ethcommon.Hash {
	slot0 := crypto.Keccak256Hash(append([]byte(poolId[:]), Six...))
	return slot0
}

func getTickSlot(poolId ethcommon.Hash, tick *big.Int) ethcommon.Hash {
	stateSlot := getStateSlot(poolId)
	four := math.NewInt(4)
	mapOffsetInt := math.NewIntFromBigInt(new(big.Int).SetBytes(stateSlot[:])).Add(four)
	tickMappingSlot := ethcommon.LeftPadBytes(mapOffsetInt.BigInt().Bytes(), 32)

	tickBz := ethcommon.LeftPadBytes(tick.Bytes(), 32)
	mapElementSlot := crypto.Keccak256Hash(append(tickBz, tickMappingSlot...))
	fmt.Println("Map element slot key:", mapElementSlot.String())
	return mapElementSlot
}

func getLiquiditySlot(poolId ethcommon.Hash) ethcommon.Hash {
	stateSlot := getStateSlot(poolId)
	liquiditySlot := math.NewIntFromBigInt(new(big.Int).SetBytes(stateSlot[:])).Add(math.NewInt(3)) // pad it by 3
	return ethcommon.BytesToHash(ethcommon.LeftPadBytes(liquiditySlot.BigInt().Bytes(), 32))
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
		denom0, denom1 = denom1, denom0
	}

	fmt.Println("denom0, 1:", denom0, denom1)
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

	fmt.Println("pool output:", res.Output)
	fmt.Println("pool info hex:", hex.EncodeToString(res.Output))
	poolInfo, err := parsePoolinfo(res.Output)
	if err != nil {
		panic(err)
	}

	bz, _ := json.Marshal(poolInfo)
	fmt.Println("poolInfo:", string(bz))

	btcPrice := 69000
	upperPrice := float64(btcPrice) * 1.2
	lowerPrice := float64(btcPrice) * 0.8
	upperTick := computeTick(upperPrice, 60)
	lowerTick := computeTick(lowerPrice, 60)
	fmt.Println("lower tick:", lowerTick.String(), "upper tick:", upperTick.String())

	// swap btc => usdt
	tickSlot := getTickSlot(poolId, lowerTick)
	tickSlotCd, err := PoolManagerABI.Pack("extsload", tickSlot)
	if err != nil {
		panic(err)
	}

	res, err = evmQc.ContractQuery(ctx, &evmtypes.ContractQueryRequest{
		Calldata: tickSlotCd,
		Address:  PoolManager,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("tick info output:", res.Output)
	// fmt.Println("tick output:", res.Output)
	// expected slot: 0x10dc5783f1a6a83fffa4decb918a065355e47495544998fde60a37f7f59ec1a3
	liquiditySlot := getLiquiditySlot(poolId)
	liquiditySlotCd, err := PoolManagerABI.Pack("extsload", liquiditySlot)
	if err != nil {
		panic(err)
	}

	fmt.Println("lq:", liquiditySlot)
	res, err = evmQc.ContractQuery(ctx, &evmtypes.ContractQueryRequest{
		Calldata: liquiditySlotCd,
		Address:  PoolManager,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("liquidity output:", res.Output)
	liquidity := math.NewIntFromBigInt(new(big.Int).SetBytes(res.Output))
	fmt.Println("liquidity number:", liquidity.String())

	poolBalanceCd, err := Erc20ABI.Pack("balanceOf", ethcommon.HexToAddress(PoolManager))
	if err != nil {
		panic(err)
	}
	res, err = evmQc.ContractQuery(ctx, &evmtypes.ContractQueryRequest{
		Calldata: poolBalanceCd,
		Address:  btcEvm,
	})
	if err != nil {
		panic(err)
	}

	balance, err := Erc20ABI.Unpack("balanceOf", res.Output)
	if err != nil {
		panic(err)
	}

	fmt.Println("btc balance in contract:", balance[0])

	poolBalanceCd, err = Erc20ABI.Pack("balanceOf", ethcommon.HexToAddress(PoolManager))
	if err != nil {
		panic(err)
	}
	res, err = evmQc.ContractQuery(ctx, &evmtypes.ContractQueryRequest{
		Calldata: poolBalanceCd,
		Address:  usdtEvm,
	})
	if err != nil {
		panic(err)
	}

	balance, err = Erc20ABI.Unpack("balanceOf", res.Output)
	if err != nil {
		panic(err)
	}

	fmt.Println("usdt balance in contract:", balance[0])
	// denom 0 = liquidity * (1/sqrt(lowerTick) - 1/sqrt(currentPrice))
	// denom 1 = liquidity * (sqrt(upperTick) - sqrt(currentPrice))
	// denom0 := liquidity.Mul(math.NewIntFromBigInt())
}

func computeSqrtPriceX96Int(p float64) *big.Int {
	sqrtPrice := new(big.Float).Sqrt(big.NewFloat(p))
	factor := new(big.Float).SetInt(new(big.Int).Lsh(big.NewInt(1), 96))
	sqrtPrice.Mul(sqrtPrice, factor)
	sqrtPriceX96Int := new(big.Int)
	price, _ := sqrtPrice.Int(sqrtPriceX96Int)
	return price
}

func logBigFloat(x *big.Float) *big.Float {
	f64, _ := x.Float64()
	return big.NewFloat(gomath.Log(f64))
}

func computeTick(price float64, spacing int64) *big.Int {
	factor := big.NewFloat(1.0001)
	priceBig := big.NewFloat(price)
	logPrice := logBigFloat(priceBig)
	logFactor := logBigFloat(factor)
	floatTick := new(big.Float).Quo(logPrice, logFactor)
	spc := big.NewInt(spacing)

	// round tick to closest spacing
	intTick := new(big.Int)
	floatTick.Int(intTick)
	roundedTick := new(big.Int).Mul(new(big.Int).Quo(intTick, spc), spc)
	return roundedTick
}

// from contract
// lower: 18565475724807095937255980671044
// upper: 22766576394583767738309774305245
// liquidity = 1000000000
// current: 20784319660459464383123105852754
// denom0: 326915
// denom1: 28349262954

// calculation
// lower = 18565475724807095937255980671044
// upper = 22766576394583767738309774305245
// current = 20811535737222300946999034249216
// l = 1000000000
// 326914
// 28349262953
// denom0 = l * (upper-current) * 2**96 // (upper*current)
// denom1 = l * (current - lower)
