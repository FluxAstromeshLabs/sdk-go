// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SettleLp is the `settleLp` instruction.
type SettleLp struct {
	MarketIndex *uint16

	// [0] = [] state
	//
	// [1] = [WRITE] user
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSettleLpInstructionBuilder creates a new `SettleLp` instruction builder.
func NewSettleLpInstructionBuilder() *SettleLp {
	nd := &SettleLp{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMarketIndex sets the "marketIndex" parameter.
func (inst *SettleLp) SetMarketIndex(marketIndex uint16) *SettleLp {
	inst.MarketIndex = &marketIndex
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *SettleLp) SetStateAccount(state ag_solanago.PublicKey) *SettleLp {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SettleLp) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserAccount sets the "user" account.
func (inst *SettleLp) SetUserAccount(user ag_solanago.PublicKey) *SettleLp {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(user).WRITE()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *SettleLp) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst SettleLp) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SettleLp,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SettleLp) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SettleLp) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MarketIndex == nil {
			return errors.New("MarketIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.User is not set")
		}
	}
	return nil
}

func (inst *SettleLp) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SettleLp")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketIndex", *inst.MarketIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" user", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj SettleLp) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketIndex` param:
	err = encoder.Encode(obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SettleLp) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketIndex`:
	err = decoder.Decode(&obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewSettleLpInstruction declares a new SettleLp instruction with the provided parameters and accounts.
func NewSettleLpInstruction(
	// Parameters:
	marketIndex uint16,
	// Accounts:
	state ag_solanago.PublicKey,
	user ag_solanago.PublicKey) *SettleLp {
	return NewSettleLpInstructionBuilder().
		SetMarketIndex(marketIndex).
		SetStateAccount(state).
		SetUserAccount(user)
}
