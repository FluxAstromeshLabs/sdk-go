// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdatePerpMarketUnrealizedAssetWeight is the `updatePerpMarketUnrealizedAssetWeight` instruction.
type UpdatePerpMarketUnrealizedAssetWeight struct {
	UnrealizedInitialAssetWeight     *uint32
	UnrealizedMaintenanceAssetWeight *uint32

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] perpMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePerpMarketUnrealizedAssetWeightInstructionBuilder creates a new `UpdatePerpMarketUnrealizedAssetWeight` instruction builder.
func NewUpdatePerpMarketUnrealizedAssetWeightInstructionBuilder() *UpdatePerpMarketUnrealizedAssetWeight {
	nd := &UpdatePerpMarketUnrealizedAssetWeight{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetUnrealizedInitialAssetWeight sets the "unrealizedInitialAssetWeight" parameter.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) SetUnrealizedInitialAssetWeight(unrealizedInitialAssetWeight uint32) *UpdatePerpMarketUnrealizedAssetWeight {
	inst.UnrealizedInitialAssetWeight = &unrealizedInitialAssetWeight
	return inst
}

// SetUnrealizedMaintenanceAssetWeight sets the "unrealizedMaintenanceAssetWeight" parameter.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) SetUnrealizedMaintenanceAssetWeight(unrealizedMaintenanceAssetWeight uint32) *UpdatePerpMarketUnrealizedAssetWeight {
	inst.UnrealizedMaintenanceAssetWeight = &unrealizedMaintenanceAssetWeight
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) SetAdminAccount(admin ag_solanago.PublicKey) *UpdatePerpMarketUnrealizedAssetWeight {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) SetStateAccount(state ag_solanago.PublicKey) *UpdatePerpMarketUnrealizedAssetWeight {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *UpdatePerpMarketUnrealizedAssetWeight {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *UpdatePerpMarketUnrealizedAssetWeight) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdatePerpMarketUnrealizedAssetWeight) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePerpMarketUnrealizedAssetWeight,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePerpMarketUnrealizedAssetWeight) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePerpMarketUnrealizedAssetWeight) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UnrealizedInitialAssetWeight == nil {
			return errors.New("UnrealizedInitialAssetWeight parameter is not set")
		}
		if inst.UnrealizedMaintenanceAssetWeight == nil {
			return errors.New("UnrealizedMaintenanceAssetWeight parameter is not set")
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

func (inst *UpdatePerpMarketUnrealizedAssetWeight) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePerpMarketUnrealizedAssetWeight")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    UnrealizedInitialAssetWeight", *inst.UnrealizedInitialAssetWeight))
						paramsBranch.Child(ag_format.Param("UnrealizedMaintenanceAssetWeight", *inst.UnrealizedMaintenanceAssetWeight))
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

func (obj UpdatePerpMarketUnrealizedAssetWeight) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UnrealizedInitialAssetWeight` param:
	err = encoder.Encode(obj.UnrealizedInitialAssetWeight)
	if err != nil {
		return err
	}
	// Serialize `UnrealizedMaintenanceAssetWeight` param:
	err = encoder.Encode(obj.UnrealizedMaintenanceAssetWeight)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdatePerpMarketUnrealizedAssetWeight) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UnrealizedInitialAssetWeight`:
	err = decoder.Decode(&obj.UnrealizedInitialAssetWeight)
	if err != nil {
		return err
	}
	// Deserialize `UnrealizedMaintenanceAssetWeight`:
	err = decoder.Decode(&obj.UnrealizedMaintenanceAssetWeight)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdatePerpMarketUnrealizedAssetWeightInstruction declares a new UpdatePerpMarketUnrealizedAssetWeight instruction with the provided parameters and accounts.
func NewUpdatePerpMarketUnrealizedAssetWeightInstruction(
	// Parameters:
	unrealizedInitialAssetWeight uint32,
	unrealizedMaintenanceAssetWeight uint32,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey) *UpdatePerpMarketUnrealizedAssetWeight {
	return NewUpdatePerpMarketUnrealizedAssetWeightInstructionBuilder().
		SetUnrealizedInitialAssetWeight(unrealizedInitialAssetWeight).
		SetUnrealizedMaintenanceAssetWeight(unrealizedMaintenanceAssetWeight).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket)
}
