// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdatePerpMarketFeeAdjustment is the `updatePerpMarketFeeAdjustment` instruction.
type UpdatePerpMarketFeeAdjustment struct {
	FeeAdjustment *int16

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] perpMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePerpMarketFeeAdjustmentInstructionBuilder creates a new `UpdatePerpMarketFeeAdjustment` instruction builder.
func NewUpdatePerpMarketFeeAdjustmentInstructionBuilder() *UpdatePerpMarketFeeAdjustment {
	nd := &UpdatePerpMarketFeeAdjustment{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetFeeAdjustment sets the "feeAdjustment" parameter.
func (inst *UpdatePerpMarketFeeAdjustment) SetFeeAdjustment(feeAdjustment int16) *UpdatePerpMarketFeeAdjustment {
	inst.FeeAdjustment = &feeAdjustment
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdatePerpMarketFeeAdjustment) SetAdminAccount(admin ag_solanago.PublicKey) *UpdatePerpMarketFeeAdjustment {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdatePerpMarketFeeAdjustment) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdatePerpMarketFeeAdjustment) SetStateAccount(state ag_solanago.PublicKey) *UpdatePerpMarketFeeAdjustment {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdatePerpMarketFeeAdjustment) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *UpdatePerpMarketFeeAdjustment) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *UpdatePerpMarketFeeAdjustment {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *UpdatePerpMarketFeeAdjustment) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdatePerpMarketFeeAdjustment) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePerpMarketFeeAdjustment,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePerpMarketFeeAdjustment) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePerpMarketFeeAdjustment) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.FeeAdjustment == nil {
			return errors.New("FeeAdjustment parameter is not set")
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

func (inst *UpdatePerpMarketFeeAdjustment) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePerpMarketFeeAdjustment")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("FeeAdjustment", *inst.FeeAdjustment))
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

func (obj UpdatePerpMarketFeeAdjustment) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `FeeAdjustment` param:
	err = encoder.Encode(obj.FeeAdjustment)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdatePerpMarketFeeAdjustment) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `FeeAdjustment`:
	err = decoder.Decode(&obj.FeeAdjustment)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdatePerpMarketFeeAdjustmentInstruction declares a new UpdatePerpMarketFeeAdjustment instruction with the provided parameters and accounts.
func NewUpdatePerpMarketFeeAdjustmentInstruction(
	// Parameters:
	feeAdjustment int16,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey) *UpdatePerpMarketFeeAdjustment {
	return NewUpdatePerpMarketFeeAdjustmentInstructionBuilder().
		SetFeeAdjustment(feeAdjustment).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket)
}