// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DeleteInitializedSpotMarket is the `deleteInitializedSpotMarket` instruction.
type DeleteInitializedSpotMarket struct {
	MarketIndex *uint16

	// [0] = [WRITE, SIGNER] admin
	//
	// [1] = [WRITE] state
	//
	// [2] = [WRITE] spotMarket
	//
	// [3] = [WRITE] spotMarketVault
	//
	// [4] = [WRITE] insuranceFundVault
	//
	// [5] = [] driftSigner
	//
	// [6] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeleteInitializedSpotMarketInstructionBuilder creates a new `DeleteInitializedSpotMarket` instruction builder.
func NewDeleteInitializedSpotMarketInstructionBuilder() *DeleteInitializedSpotMarket {
	nd := &DeleteInitializedSpotMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetMarketIndex sets the "marketIndex" parameter.
func (inst *DeleteInitializedSpotMarket) SetMarketIndex(marketIndex uint16) *DeleteInitializedSpotMarket {
	inst.MarketIndex = &marketIndex
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *DeleteInitializedSpotMarket) SetAdminAccount(admin ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *DeleteInitializedSpotMarket) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *DeleteInitializedSpotMarket) SetStateAccount(state ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *DeleteInitializedSpotMarket) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *DeleteInitializedSpotMarket) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *DeleteInitializedSpotMarket) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSpotMarketVaultAccount sets the "spotMarketVault" account.
func (inst *DeleteInitializedSpotMarket) SetSpotMarketVaultAccount(spotMarketVault ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(spotMarketVault).WRITE()
	return inst
}

// GetSpotMarketVaultAccount gets the "spotMarketVault" account.
func (inst *DeleteInitializedSpotMarket) GetSpotMarketVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetInsuranceFundVaultAccount sets the "insuranceFundVault" account.
func (inst *DeleteInitializedSpotMarket) SetInsuranceFundVaultAccount(insuranceFundVault ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(insuranceFundVault).WRITE()
	return inst
}

// GetInsuranceFundVaultAccount gets the "insuranceFundVault" account.
func (inst *DeleteInitializedSpotMarket) GetInsuranceFundVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetDriftSignerAccount sets the "driftSigner" account.
func (inst *DeleteInitializedSpotMarket) SetDriftSignerAccount(driftSigner ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(driftSigner)
	return inst
}

// GetDriftSignerAccount gets the "driftSigner" account.
func (inst *DeleteInitializedSpotMarket) GetDriftSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *DeleteInitializedSpotMarket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *DeleteInitializedSpotMarket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst DeleteInitializedSpotMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DeleteInitializedSpotMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeleteInitializedSpotMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeleteInitializedSpotMarket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MarketIndex == nil {
			return errors.New("MarketIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SpotMarket is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SpotMarketVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.InsuranceFundVault is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.DriftSigner is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *DeleteInitializedSpotMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeleteInitializedSpotMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketIndex", *inst.MarketIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             state", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        spotMarket", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   spotMarketVault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("insuranceFundVault", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       driftSigner", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj DeleteInitializedSpotMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketIndex` param:
	err = encoder.Encode(obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DeleteInitializedSpotMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketIndex`:
	err = decoder.Decode(&obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewDeleteInitializedSpotMarketInstruction declares a new DeleteInitializedSpotMarket instruction with the provided parameters and accounts.
func NewDeleteInitializedSpotMarketInstruction(
	// Parameters:
	marketIndex uint16,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	spotMarket ag_solanago.PublicKey,
	spotMarketVault ag_solanago.PublicKey,
	insuranceFundVault ag_solanago.PublicKey,
	driftSigner ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *DeleteInitializedSpotMarket {
	return NewDeleteInitializedSpotMarketInstructionBuilder().
		SetMarketIndex(marketIndex).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetSpotMarketAccount(spotMarket).
		SetSpotMarketVaultAccount(spotMarketVault).
		SetInsuranceFundVaultAccount(insuranceFundVault).
		SetDriftSignerAccount(driftSigner).
		SetTokenProgramAccount(tokenProgram)
}