// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateSpotMarketLiquidationFee is the `updateSpotMarketLiquidationFee` instruction.
type UpdateSpotMarketLiquidationFee struct {
	LiquidatorFee    *uint32
	IfLiquidationFee *uint32

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] spotMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateSpotMarketLiquidationFeeInstructionBuilder creates a new `UpdateSpotMarketLiquidationFee` instruction builder.
func NewUpdateSpotMarketLiquidationFeeInstructionBuilder() *UpdateSpotMarketLiquidationFee {
	nd := &UpdateSpotMarketLiquidationFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetLiquidatorFee sets the "liquidatorFee" parameter.
func (inst *UpdateSpotMarketLiquidationFee) SetLiquidatorFee(liquidatorFee uint32) *UpdateSpotMarketLiquidationFee {
	inst.LiquidatorFee = &liquidatorFee
	return inst
}

// SetIfLiquidationFee sets the "ifLiquidationFee" parameter.
func (inst *UpdateSpotMarketLiquidationFee) SetIfLiquidationFee(ifLiquidationFee uint32) *UpdateSpotMarketLiquidationFee {
	inst.IfLiquidationFee = &ifLiquidationFee
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateSpotMarketLiquidationFee) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateSpotMarketLiquidationFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateSpotMarketLiquidationFee) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdateSpotMarketLiquidationFee) SetStateAccount(state ag_solanago.PublicKey) *UpdateSpotMarketLiquidationFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateSpotMarketLiquidationFee) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *UpdateSpotMarketLiquidationFee) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *UpdateSpotMarketLiquidationFee {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *UpdateSpotMarketLiquidationFee) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateSpotMarketLiquidationFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateSpotMarketLiquidationFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateSpotMarketLiquidationFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateSpotMarketLiquidationFee) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.LiquidatorFee == nil {
			return errors.New("LiquidatorFee parameter is not set")
		}
		if inst.IfLiquidationFee == nil {
			return errors.New("IfLiquidationFee parameter is not set")
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
	}
	return nil
}

func (inst *UpdateSpotMarketLiquidationFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateSpotMarketLiquidationFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   LiquidatorFee", *inst.LiquidatorFee))
						paramsBranch.Child(ag_format.Param("IfLiquidationFee", *inst.IfLiquidationFee))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     state", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("spotMarket", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdateSpotMarketLiquidationFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LiquidatorFee` param:
	err = encoder.Encode(obj.LiquidatorFee)
	if err != nil {
		return err
	}
	// Serialize `IfLiquidationFee` param:
	err = encoder.Encode(obj.IfLiquidationFee)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateSpotMarketLiquidationFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LiquidatorFee`:
	err = decoder.Decode(&obj.LiquidatorFee)
	if err != nil {
		return err
	}
	// Deserialize `IfLiquidationFee`:
	err = decoder.Decode(&obj.IfLiquidationFee)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateSpotMarketLiquidationFeeInstruction declares a new UpdateSpotMarketLiquidationFee instruction with the provided parameters and accounts.
func NewUpdateSpotMarketLiquidationFeeInstruction(
	// Parameters:
	liquidatorFee uint32,
	ifLiquidationFee uint32,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	spotMarket ag_solanago.PublicKey) *UpdateSpotMarketLiquidationFee {
	return NewUpdateSpotMarketLiquidationFeeInstructionBuilder().
		SetLiquidatorFee(liquidatorFee).
		SetIfLiquidationFee(ifLiquidationFee).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetSpotMarketAccount(spotMarket)
}