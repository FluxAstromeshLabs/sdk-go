// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateSpotMarketName is the `updateSpotMarketName` instruction.
type UpdateSpotMarketName struct {
	Name *[32]uint8

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] spotMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateSpotMarketNameInstructionBuilder creates a new `UpdateSpotMarketName` instruction builder.
func NewUpdateSpotMarketNameInstructionBuilder() *UpdateSpotMarketName {
	nd := &UpdateSpotMarketName{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetName sets the "name" parameter.
func (inst *UpdateSpotMarketName) SetName(name [32]uint8) *UpdateSpotMarketName {
	inst.Name = &name
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateSpotMarketName) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateSpotMarketName {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateSpotMarketName) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdateSpotMarketName) SetStateAccount(state ag_solanago.PublicKey) *UpdateSpotMarketName {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateSpotMarketName) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *UpdateSpotMarketName) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *UpdateSpotMarketName {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *UpdateSpotMarketName) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateSpotMarketName) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateSpotMarketName,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateSpotMarketName) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateSpotMarketName) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
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

func (inst *UpdateSpotMarketName) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateSpotMarketName")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Name", *inst.Name))
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

func (obj UpdateSpotMarketName) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateSpotMarketName) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateSpotMarketNameInstruction declares a new UpdateSpotMarketName instruction with the provided parameters and accounts.
func NewUpdateSpotMarketNameInstruction(
	// Parameters:
	name [32]uint8,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	spotMarket ag_solanago.PublicKey) *UpdateSpotMarketName {
	return NewUpdateSpotMarketNameInstructionBuilder().
		SetName(name).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetSpotMarketAccount(spotMarket)
}