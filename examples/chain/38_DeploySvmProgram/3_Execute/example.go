package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	sdkmath "cosmossdk.io/math"
	svmtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/svm/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/FluxNFTLabs/sdk-go/client/svm"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ethsecp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BuildSignedTx(
	chainClient chainclient.ChainClient,
	msgs []sdk.Msg,
	cosmosSignerKeys []*ethsecp256k1.PrivKey,
) (sdk.Tx, error) {
	cosmosAccs := []sdk.AccAddress{}
	cosmosPubkeys := []cryptotypes.PubKey{}
	for _, signer := range cosmosSignerKeys {
		cosmosPubkeys = append(cosmosPubkeys, signer.PubKey())
		acc := sdk.AccAddress(signer.PubKey().Address().Bytes())
		cosmosAccs = append(cosmosAccs, acc)
	}
	userNums := make([]uint64, len(cosmosSignerKeys))
	userSeqs := make([]uint64, len(cosmosSignerKeys))
	clientCtx := chainClient.ClientContext()
	// init tx builder
	for i, userAddr := range cosmosAccs {
		userNum, userSeq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(clientCtx, userAddr)
		if err != nil {
			return nil, fmt.Errorf("get acc number err: %w", err)
		}

		userNums[i] = userNum
		userSeqs[i] = userSeq
	}

	txBuilder := clientCtx.TxConfig.NewTxBuilder()
	// prepare tx data
	timeoutHeight := uint64(19000000)
	gasLimit := uint64(30000000)
	gasPrice := sdkmath.NewIntFromUint64(500000000)
	fee := sdk.NewCoins(sdk.NewCoin("lux", sdkmath.NewIntFromUint64(gasLimit).Mul(gasPrice)))
	signatures := make([]signingtypes.SignatureV2, len(cosmosSignerKeys))

	for i := range cosmosPubkeys {
		userSigV2 := signingtypes.SignatureV2{
			PubKey: cosmosPubkeys[i],
			Data: &signingtypes.SingleSignatureData{
				SignMode:  signingtypes.SignMode_SIGN_MODE_DIRECT,
				Signature: nil,
			},
			Sequence: userSeqs[i],
		}
		signatures[i] = userSigV2
	}

	// build tx
	txBuilder.SetMsgs(msgs...)
	txBuilder.SetGasLimit(gasLimit)
	txBuilder.SetTimeoutHeight(timeoutHeight)
	txBuilder.SetFeeAmount(fee)
	txBuilder.SetSignatures(signatures...)

	// build and sign tx
	for i, pk := range cosmosSignerKeys {
		sig, err := tx.SignWithPrivKey(
			context.Background(),
			signingtypes.SignMode_SIGN_MODE_DIRECT,
			authsigning.SignerData{
				Address:       cosmosAccs[i].String(),
				ChainID:       clientCtx.ChainID,
				AccountNumber: userNums[i],
				Sequence:      userSeqs[i],
				PubKey:        pk.PubKey(),
			},
			txBuilder, pk, clientCtx.TxConfig, userSeqs[i],
		)
		if err != nil {
			panic(err)
		}

		signatures[i] = sig
	}

	// set signatures again
	txBuilder.SetSignatures(signatures...)
	// build sign data
	return txBuilder.GetTx(), nil
}

func LinkAccount(
	chainClient chainclient.ChainClient,
	clientCtx client.Context,
	cosmosPrivKey *ethsecp256k1.PrivKey,
	svmPrivKey *ed25519.PrivKey,
	luxAmount int64,
) (*txtypes.BroadcastTxResponse, error) {
	// prepare users
	userPubKey := cosmosPrivKey.PubKey()
	userAddr := sdk.AccAddress(userPubKey.Address().Bytes())

	// init link msg
	svmPubkey := svmPrivKey.PubKey()
	svmSig, err := svmPrivKey.Sign(userAddr.Bytes())
	if err != nil {
		return nil, fmt.Errorf("svm private key sign err: %w", err)
	}

	msg := &svmtypes.MsgLinkSVMAccount{
		Sender:       userAddr.String(),
		SvmPubkey:    svmPubkey.Bytes(),
		SvmSignature: svmSig[:],
		Amount:       sdk.NewInt64Coin("lux", luxAmount),
	}

	signedTx, err := BuildSignedTx(chainClient, []sdk.Msg{msg}, []*ethsecp256k1.PrivKey{cosmosPrivKey})
	if err != nil {
		return nil, err
	}

	txBytes, err := chainClient.ClientContext().TxConfig.TxEncoder()(signedTx)
	if err != nil {
		return nil, err
	}

	return chainClient.SyncBroadcastSignedTx(txBytes)
}

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
		common.OptionGasPrices("500000000lux"),
	)
	if err != nil {
		panic(err)
	}

	cosmosPrivateKeys := []*ethsecp256k1.PrivKey{
		{Key: ethcommon.Hex2Bytes("88cbead91aee890d27bf06e003ade3d4e952427e88f88d31d61d3ef5e5d54305")},
		{Key: ethcommon.Hex2Bytes("c25e5cccd433d2c97971eaa6cfe92ea05771dc05b984c62464ab580f16a905e1")},
	}
	cosmosAddrs := make([]sdk.AccAddress, len(cosmosPrivateKeys))
	for i, pk := range cosmosPrivateKeys {
		cosmosAddrs[i] = sdk.AccAddress(pk.PubKey().Address().Bytes())
	}

	programOwnerSvmPrivKey := ed25519.GenPrivKeyFromSecret([]byte("owner"))
	programOwnerPubkey := solana.PublicKeyFromBytes(programOwnerSvmPrivKey.PubKey().Bytes())

	// when executing, we only know the program pubkey
	programPubkey := solana.MustPublicKeyFromBase58("Eutyt8ETuT9Zsjj8z6hxdwnzaoQadgMzANr25t3DyMk3")
	programDataSvmPrivKey := ed25519.GenPrivKeyFromSecret([]byte("helloProgramData"))
	programDataSvmPubKey := solana.PublicKeyFromBytes(programDataSvmPrivKey.PubKey().Bytes())
	programDataCosmosAddr := cosmosAddrs[1]

	isSvmLinked, _, err := chainClient.GetSVMAccountLink(context.Background(), programDataCosmosAddr)
	if err != nil {
		panic(err)
	}

	if !isSvmLinked {
		fmt.Printf("new cosmos account %s is unliked to svm. Going to link it now\n", programDataCosmosAddr.String())
		_, err = LinkAccount(chainClient, clientCtx, cosmosPrivateKeys[1], programDataSvmPrivKey, 0)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("linked %s to svm %s\n", programDataCosmosAddr.String(), programDataSvmPubKey.String())
	tx := solana.NewTransactionBuilder()
	accRes, err := chainClient.GetSvmAccount(context.Background(), programDataSvmPubKey.String())
	if err != nil {
		panic(err)
	}
	// fund some lamports if it's zero
	if accRes.Account.Lamports == 0 {
		accountSz := uint64(4)
		createAccountIx := system.NewCreateAccountInstruction(
			svmtypes.GetRentExemptLamportAmount(accountSz), accountSz,
			programPubkey,
			programOwnerPubkey,
			programDataSvmPubKey,
		).Build()
		tx = tx.AddInstruction(createAccountIx)
	}

	// execute program ix
	tx.AddInstruction(solana.NewInstruction(
		programPubkey,
		solana.AccountMetaSlice{
			{
				PublicKey:  programDataSvmPubKey,
				IsWritable: true,
				IsSigner:   true,
			},
		},
		nil,
	))

	executeTx, err := tx.Build()
	if err != nil {
		panic(err)
	}

	signers := []string{}
	for _, acc := range cosmosAddrs {
		signers = append(signers, acc.String())
	}

	cosmosMsg := svm.ToCosmosMsg(signers, 10_000_000, executeTx)
	signedTx, err := BuildSignedTx(chainClient, []sdk.Msg{cosmosMsg}, cosmosPrivateKeys)
	if err != nil {
		panic(err)
	}

	txBytes, err := chainClient.ClientContext().TxConfig.TxEncoder()(signedTx)
	if err != nil {
		panic(err)
	}

	res, err := chainClient.SyncBroadcastSignedTx(txBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("tx hash:", res.TxResponse.TxHash)
	fmt.Println("gas used/want:", res.TxResponse.GasUsed, "/", res.TxResponse.GasWanted)
}
