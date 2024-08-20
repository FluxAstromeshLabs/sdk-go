// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateInsuranceFundUnstakingPeriod is the `updateInsuranceFundUnstakingPeriod` instruction.
type UpdateInsuranceFundUnstakingPeriod struct {
	InsuranceFundUnstakingPeriod *int64

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] spotMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateInsuranceFundUnstakingPeriodInstructionBuilder creates a new `UpdateInsuranceFundUnstakingPeriod` instruction builder.
func NewUpdateInsuranceFundUnstakingPeriodInstructionBuilder() *UpdateInsuranceFundUnstakingPeriod {
	nd := &UpdateInsuranceFundUnstakingPeriod{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetInsuranceFundUnstakingPeriod sets the "insuranceFundUnstakingPeriod" parameter.
func (inst *UpdateInsuranceFundUnstakingPeriod) SetInsuranceFundUnstakingPeriod(insuranceFundUnstakingPeriod int64) *UpdateInsuranceFundUnstakingPeriod {
	inst.InsuranceFundUnstakingPeriod = &insuranceFundUnstakingPeriod
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateInsuranceFundUnstakingPeriod) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateInsuranceFundUnstakingPeriod {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateInsuranceFundUnstakingPeriod) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdateInsuranceFundUnstakingPeriod) SetStateAccount(state ag_solanago.PublicKey) *UpdateInsuranceFundUnstakingPeriod {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateInsuranceFundUnstakingPeriod) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *UpdateInsuranceFundUnstakingPeriod) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *UpdateInsuranceFundUnstakingPeriod {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *UpdateInsuranceFundUnstakingPeriod) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateInsuranceFundUnstakingPeriod) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateInsuranceFundUnstakingPeriod,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateInsuranceFundUnstakingPeriod) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateInsuranceFundUnstakingPeriod) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.InsuranceFundUnstakingPeriod == nil {
			return errors.New("InsuranceFundUnstakingPeriod parameter is not set")
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

func (inst *UpdateInsuranceFundUnstakingPeriod) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateInsuranceFundUnstakingPeriod")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("InsuranceFundUnstakingPeriod", *inst.InsuranceFundUnstakingPeriod))
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

func (obj UpdateInsuranceFundUnstakingPeriod) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `InsuranceFundUnstakingPeriod` param:
	err = encoder.Encode(obj.InsuranceFundUnstakingPeriod)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateInsuranceFundUnstakingPeriod) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `InsuranceFundUnstakingPeriod`:
	err = decoder.Decode(&obj.InsuranceFundUnstakingPeriod)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateInsuranceFundUnstakingPeriodInstruction declares a new UpdateInsuranceFundUnstakingPeriod instruction with the provided parameters and accounts.
func NewUpdateInsuranceFundUnstakingPeriodInstruction(
	// Parameters:
	insuranceFundUnstakingPeriod int64,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	spotMarket ag_solanago.PublicKey) *UpdateInsuranceFundUnstakingPeriod {
	return NewUpdateInsuranceFundUnstakingPeriodInstructionBuilder().
		SetInsuranceFundUnstakingPeriod(insuranceFundUnstakingPeriod).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetSpotMarketAccount(spotMarket)
}
