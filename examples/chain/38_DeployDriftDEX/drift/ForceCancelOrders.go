// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ForceCancelOrders is the `forceCancelOrders` instruction.
type ForceCancelOrders struct {

	// [0] = [] state
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [WRITE] filler
	//
	// [3] = [WRITE] user
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewForceCancelOrdersInstructionBuilder creates a new `ForceCancelOrders` instruction builder.
func NewForceCancelOrdersInstructionBuilder() *ForceCancelOrders {
	nd := &ForceCancelOrders{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *ForceCancelOrders) SetStateAccount(state ag_solanago.PublicKey) *ForceCancelOrders {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *ForceCancelOrders) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ForceCancelOrders) SetAuthorityAccount(authority ag_solanago.PublicKey) *ForceCancelOrders {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ForceCancelOrders) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFillerAccount sets the "filler" account.
func (inst *ForceCancelOrders) SetFillerAccount(filler ag_solanago.PublicKey) *ForceCancelOrders {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(filler).WRITE()
	return inst
}

// GetFillerAccount gets the "filler" account.
func (inst *ForceCancelOrders) GetFillerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserAccount sets the "user" account.
func (inst *ForceCancelOrders) SetUserAccount(user ag_solanago.PublicKey) *ForceCancelOrders {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(user).WRITE()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *ForceCancelOrders) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst ForceCancelOrders) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ForceCancelOrders,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ForceCancelOrders) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ForceCancelOrders) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Filler is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.User is not set")
		}
	}
	return nil
}

func (inst *ForceCancelOrders) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ForceCancelOrders")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   filler", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     user", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj ForceCancelOrders) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ForceCancelOrders) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewForceCancelOrdersInstruction declares a new ForceCancelOrders instruction with the provided parameters and accounts.
func NewForceCancelOrdersInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	filler ag_solanago.PublicKey,
	user ag_solanago.PublicKey) *ForceCancelOrders {
	return NewForceCancelOrdersInstructionBuilder().
		SetStateAccount(state).
		SetAuthorityAccount(authority).
		SetFillerAccount(filler).
		SetUserAccount(user)
}