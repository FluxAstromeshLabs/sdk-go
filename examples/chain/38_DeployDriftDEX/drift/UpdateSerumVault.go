// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateSerumVault is the `updateSerumVault` instruction.
type UpdateSerumVault struct {

	// [0] = [WRITE] state
	//
	// [1] = [WRITE, SIGNER] admin
	//
	// [2] = [] srmVault
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateSerumVaultInstructionBuilder creates a new `UpdateSerumVault` instruction builder.
func NewUpdateSerumVaultInstructionBuilder() *UpdateSerumVault {
	nd := &UpdateSerumVault{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *UpdateSerumVault) SetStateAccount(state ag_solanago.PublicKey) *UpdateSerumVault {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateSerumVault) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateSerumVault) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateSerumVault {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateSerumVault) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSrmVaultAccount sets the "srmVault" account.
func (inst *UpdateSerumVault) SetSrmVaultAccount(srmVault ag_solanago.PublicKey) *UpdateSerumVault {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(srmVault)
	return inst
}

// GetSrmVaultAccount gets the "srmVault" account.
func (inst *UpdateSerumVault) GetSrmVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateSerumVault) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateSerumVault,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateSerumVault) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateSerumVault) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SrmVault is not set")
		}
	}
	return nil
}

func (inst *UpdateSerumVault) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateSerumVault")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("srmVault", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdateSerumVault) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UpdateSerumVault) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUpdateSerumVaultInstruction declares a new UpdateSerumVault instruction with the provided parameters and accounts.
func NewUpdateSerumVaultInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	srmVault ag_solanago.PublicKey) *UpdateSerumVault {
	return NewUpdateSerumVaultInstructionBuilder().
		SetStateAccount(state).
		SetAdminAccount(admin).
		SetSrmVaultAccount(srmVault)
}
