// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdatePerpMarketPausedOperations is the `updatePerpMarketPausedOperations` instruction.
type UpdatePerpMarketPausedOperations struct {
	PausedOperations *uint8

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] perpMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePerpMarketPausedOperationsInstructionBuilder creates a new `UpdatePerpMarketPausedOperations` instruction builder.
func NewUpdatePerpMarketPausedOperationsInstructionBuilder() *UpdatePerpMarketPausedOperations {
	nd := &UpdatePerpMarketPausedOperations{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetPausedOperations sets the "pausedOperations" parameter.
func (inst *UpdatePerpMarketPausedOperations) SetPausedOperations(pausedOperations uint8) *UpdatePerpMarketPausedOperations {
	inst.PausedOperations = &pausedOperations
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdatePerpMarketPausedOperations) SetAdminAccount(admin ag_solanago.PublicKey) *UpdatePerpMarketPausedOperations {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdatePerpMarketPausedOperations) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdatePerpMarketPausedOperations) SetStateAccount(state ag_solanago.PublicKey) *UpdatePerpMarketPausedOperations {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdatePerpMarketPausedOperations) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *UpdatePerpMarketPausedOperations) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *UpdatePerpMarketPausedOperations {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *UpdatePerpMarketPausedOperations) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdatePerpMarketPausedOperations) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePerpMarketPausedOperations,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePerpMarketPausedOperations) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePerpMarketPausedOperations) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PausedOperations == nil {
			return errors.New("PausedOperations parameter is not set")
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
			return errors.New("accounts.PerpMarket is not set")
		}
	}
	return nil
}

func (inst *UpdatePerpMarketPausedOperations) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePerpMarketPausedOperations")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("PausedOperations", *inst.PausedOperations))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     state", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("perpMarket", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdatePerpMarketPausedOperations) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PausedOperations` param:
	err = encoder.Encode(obj.PausedOperations)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdatePerpMarketPausedOperations) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PausedOperations`:
	err = decoder.Decode(&obj.PausedOperations)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdatePerpMarketPausedOperationsInstruction declares a new UpdatePerpMarketPausedOperations instruction with the provided parameters and accounts.
func NewUpdatePerpMarketPausedOperationsInstruction(
	// Parameters:
	pausedOperations uint8,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey) *UpdatePerpMarketPausedOperations {
	return NewUpdatePerpMarketPausedOperationsInstructionBuilder().
		SetPausedOperations(pausedOperations).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket)
}