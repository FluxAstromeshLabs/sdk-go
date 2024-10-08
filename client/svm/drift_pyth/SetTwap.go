// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package pyth

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetTwap is the `setTwap` instruction.
type SetTwap struct {
	Twap *int64

	// [0] = [WRITE] price
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetTwapInstructionBuilder creates a new `SetTwap` instruction builder.
func NewSetTwapInstructionBuilder() *SetTwap {
	nd := &SetTwap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetTwap sets the "twap" parameter.
func (inst *SetTwap) SetTwap(twap int64) *SetTwap {
	inst.Twap = &twap
	return inst
}

// SetPriceAccount sets the "price" account.
func (inst *SetTwap) SetPriceAccount(price ag_solanago.PublicKey) *SetTwap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(price).WRITE()
	return inst
}

// GetPriceAccount gets the "price" account.
func (inst *SetTwap) GetPriceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

func (inst SetTwap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetTwap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetTwap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetTwap) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Twap == nil {
			return errors.New("Twap parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Price is not set")
		}
	}
	return nil
}

func (inst *SetTwap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetTwap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Twap", *inst.Twap))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("price", inst.AccountMetaSlice.Get(0)))
					})
				})
		})
}

func (obj SetTwap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Twap` param:
	err = encoder.Encode(obj.Twap)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetTwap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Twap`:
	err = decoder.Decode(&obj.Twap)
	if err != nil {
		return err
	}
	return nil
}

// NewSetTwapInstruction declares a new SetTwap instruction with the provided parameters and accounts.
func NewSetTwapInstruction(
	// Parameters:
	twap int64,
	// Accounts:
	price ag_solanago.PublicKey) *SetTwap {
	return NewSetTwapInstructionBuilder().
		SetTwap(twap).
		SetPriceAccount(price)
}
