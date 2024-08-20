// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializePredictionMarket is the `initializePredictionMarket` instruction.
type InitializePredictionMarket struct {

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] perpMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializePredictionMarketInstructionBuilder creates a new `InitializePredictionMarket` instruction builder.
func NewInitializePredictionMarketInstructionBuilder() *InitializePredictionMarket {
	nd := &InitializePredictionMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAdminAccount sets the "admin" account.
func (inst *InitializePredictionMarket) SetAdminAccount(admin ag_solanago.PublicKey) *InitializePredictionMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitializePredictionMarket) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *InitializePredictionMarket) SetStateAccount(state ag_solanago.PublicKey) *InitializePredictionMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializePredictionMarket) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *InitializePredictionMarket) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *InitializePredictionMarket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *InitializePredictionMarket) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst InitializePredictionMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializePredictionMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializePredictionMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializePredictionMarket) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PerpMarket is not set")
		}
	}
	return nil
}

func (inst *InitializePredictionMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializePredictionMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     state", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("perpMarket", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj InitializePredictionMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitializePredictionMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializePredictionMarketInstruction declares a new InitializePredictionMarket instruction with the provided parameters and accounts.
func NewInitializePredictionMarketInstruction(
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey) *InitializePredictionMarket {
	return NewInitializePredictionMarketInstructionBuilder().
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket)
}
