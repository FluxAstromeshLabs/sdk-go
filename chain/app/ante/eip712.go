package ante

import (
	"context"
	"fmt"
	"strconv"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txsigning "cosmossdk.io/x/tx/signing"

	log "github.com/InjectiveLabs/suplog"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ethsecp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ethlibsecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/FluxNFTLabs/sdk-go/chain/app/ante/typeddata"

	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
)

// Verify all signatures for a tx and return an error if any are invalid. Note,
// the Eip712SigVerificationDecorator decorator will not get executed on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type Eip712SigVerificationDecorator struct {
	ak              AccountKeeper
	signModeHandler *txsigning.HandlerMap
}

func NewEip712SigVerificationDecorator(ak AccountKeeper, signModeHandler *txsigning.HandlerMap) Eip712SigVerificationDecorator {
	return Eip712SigVerificationDecorator{
		ak:              ak,
		signModeHandler: signModeHandler,
	}
}

func (svd Eip712SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// no need to verify signatures on recheck tx
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	signerAddrs, err := sigTx.GetSigners()
	if err != nil {
		return ctx, err
	}

	// check that signer length and signature length are the same
	if len(sigs) != len(signerAddrs) {
		return ctx, errors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
	}

	for i, sig := range sigs {
		acc, err := authante.GetSignerAcc(ctx, svd.ak, signerAddrs[i])
		if err != nil {
			return ctx, err
		}

		// retrieve pubkey
		pubKey := acc.GetPubKey()
		if !simulate && pubKey == nil {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
		}

		// Check account sequence number.
		if sig.Sequence != acc.GetSequence() {
			return ctx, errors.Wrapf(
				sdkerrors.ErrWrongSequence,
				"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
			)
		}

		// retrieve signer data
		genesis := ctx.BlockHeight() == 0
		chainID := ctx.ChainID()
		var accNum uint64
		if !genesis {
			accNum = acc.GetAccountNumber()
		}
		signerData := txsigning.SignerData{
			ChainID:       chainID,
			Address:       acc.GetAddress().String(),
			AccountNumber: accNum,
			Sequence:      acc.GetSequence(),
		}

		if !simulate {
			err := VerifySignatureEIP712(pubKey, signerData, sig.Data, svd.signModeHandler, tx.(authsigning.Tx))
			if err != nil {
				log.WithError(err).Errorln("Eip712SigVerificationDecorator failed to verify signature")

				errMsg := fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s)", accNum, chainID)
				return ctx, errors.Wrap(sdkerrors.ErrUnauthorized, errMsg)

			}
		}
	}

	return next(ctx, tx, simulate)
}

var chainTypesCodec codec.ProtoCodecMarshaler

func init() {
	registry := codectypes.NewInterfaceRegistry()
	chaintypes.RegisterInterfaces(registry)
	chainTypesCodec = codec.NewProtoCodec(registry)
}

// VerifySignature verifies a transaction signature contained in SignatureData abstracting over different signing modes
// and single vs multi-signatures.
func VerifySignatureEIP712(
	pubKey cryptotypes.PubKey,
	signerData txsigning.SignerData,
	sigData signing.SignatureData,
	handler *txsigning.HandlerMap,
	tx authsigning.Tx,
) error {
	switch data := sigData.(type) {
	case *signing.SingleSignatureData:
		if data.SignMode != signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON {
			return fmt.Errorf("unexpected SignatureData %T: wrong SignMode", sigData)
		}

		// tweak signing tx data to remove extension option here to not break aminojson impl
		signingTxData := tx.(authsigning.V2AdaptableTx).GetSigningTxData()
		signingTxData.Body.ExtensionOptions = nil
		signingTxData.Body.NonCriticalExtensionOptions = nil
		msgs := tx.GetMsgs()
		txBytes, err := handler.GetSignBytes(
			context.Background(),
			signingv1beta1.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
			signerData,
			signingTxData,
		)
		if err != nil {
			return err
		}

		var (
			chainID      uint64
			feePayer     sdk.AccAddress
			feePayerSig  []byte
			feeDelegated bool
		)

		if txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx); ok {
			if opts := txWithExtensions.GetExtensionOptions(); len(opts) > 0 {
				var optIface txtypes.TxExtensionOptionI
				if err := chainTypesCodec.UnpackAny(opts[0], &optIface); err != nil {
					err = errors.Wrap(err, "failed to proto-unpack ExtensionOptionsWeb3Tx")
					return err
				} else if extOpt, ok := optIface.(*chaintypes.ExtensionOptionsWeb3Tx); ok {
					// chainID in EIP712 typed data is allowed to not match signerData.ChainID,
					// but limited to certain options: 1 (mainnet), 5 (Goerli), thus Metamask will
					// be able to submit signatures without switching networks.
					if extOpt.TypedDataChainID == 1 || extOpt.TypedDataChainID == 5 {
						chainID = extOpt.TypedDataChainID
					}

					if len(extOpt.FeePayer) > 0 {
						feePayer, err = sdk.AccAddressFromBech32(extOpt.FeePayer)
						if err != nil {
							err = errors.Wrap(err, "failed to parse feePayer from ExtensionOptionsWeb3Tx")
							return err
						}

						feePayerSig = extOpt.FeePayerSig
						if len(feePayerSig) == 0 {
							return fmt.Errorf("no feePayerSig provided in ExtensionOptionsWeb3Tx")
						}

						feeDelegated = true
					}
				}
			}
		}

		if chainID == 0 {
			chainID, err = strconv.ParseUint(signerData.ChainID, 10, 64)
			if err != nil {
				err = errors.Wrapf(err, "failed to parse chainID: %s", signerData.ChainID)
				return err
			}
		}

		var typedData typeddata.TypedData
		var sigHash []byte

		if feeDelegated {
			feeDelegation := &FeeDelegationOptions{
				FeePayer: feePayer,
			}

			typedData, err = WrapTxToEIP712(chainTypesCodec, chainID, msgs[0], txBytes, feeDelegation)
			if err != nil {
				err = errors.Wrap(err, "failed to pack tx data in EIP712 object")
				return err
			}

			sigHash, err = typeddata.ComputeTypedDataHash(typedData)
			if err != nil {
				return err
			}

			feePayerPubkey, err := ethlibsecp256k1.RecoverPubkey(sigHash, feePayerSig)
			if err != nil {
				err = errors.Wrap(err, "failed to recover delegated fee payer from sig")
				return err
			}

			ecPubKey, err := ethcrypto.UnmarshalPubkey(feePayerPubkey)
			if err != nil {
				err = errors.Wrap(err, "failed to unmarshal recovered fee payer pubkey")
				return err
			}

			recoveredFeePayerAcc := sdk.AccAddress((&ethsecp256k1.PubKey{
				Key: ethcrypto.CompressPubkey(ecPubKey),
			}).Address().Bytes())

			if !recoveredFeePayerAcc.Equals(feePayer) {
				err = fmt.Errorf("failed to verify delegated fee payer sig")
				return err
			}

		} else {
			typedData, err = WrapTxToEIP712(chainTypesCodec, chainID, msgs[0], txBytes, nil)
			if err != nil {
				err = errors.Wrap(err, "failed to pack tx data in EIP712 object")
				return err
			}

			sigHash, err = typeddata.ComputeTypedDataHash(typedData)
			if err != nil {
				return err
			}
		}

		if len(data.Signature) != 65 {
			return fmt.Errorf("signature length doesnt match typical [R||S||V] signature 65 bytes")
		}

		// VerifySignature of ethsecp256k1 accepts 64 byte signature [R||S]
		// WARNING! Under NO CIRCUMSTANCES try to use pubKey.VerifySignature there
		if !ethlibsecp256k1.VerifySignature(pubKey.Bytes(), sigHash, data.Signature[:len(data.Signature)-1]) {
			return fmt.Errorf("unable to verify signer signature of EIP712 typed data")
		}

		return nil
	default:
		return fmt.Errorf("unexpected SignatureData %T", sigData)
	}
}
