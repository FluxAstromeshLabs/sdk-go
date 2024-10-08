// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package pyth

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetPrice is the `setPrice` instruction.
type SetPrice struct {
	Price *int64

	// [0] = [WRITE] price
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetPriceInstructionBuilder creates a new `SetPrice` instruction builder.
func NewSetPriceInstructionBuilder() *SetPrice {
	nd := &SetPrice{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetPrice sets the "price" parameter.
func (inst *SetPrice) SetPrice(price int64) *SetPrice {
	inst.Price = &price
	return inst
}

// SetPriceAccount sets the "price" account.
func (inst *SetPrice) SetPriceAccount(price ag_solanago.PublicKey) *SetPrice {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(price).WRITE()
	return inst
}

// GetPriceAccount gets the "price" account.
func (inst *SetPrice) GetPriceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

func (inst SetPrice) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetPrice,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetPrice) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetPrice) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Price == nil {
			return errors.New("Price parameter is not set")
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

func (inst *SetPrice) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetPrice")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Price", *inst.Price))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("price", inst.AccountMetaSlice.Get(0)))
					})
				})
		})
}

func (obj SetPrice) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetPrice) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	return nil
}

// NewSetPriceInstruction declares a new SetPrice instruction with the provided parameters and accounts.
func NewSetPriceInstruction(
	// Parameters:
	price int64,
	// Accounts:
	priceAccount ag_solanago.PublicKey) *SetPrice {
	return NewSetPriceInstructionBuilder().
		SetPrice(price).
		SetPriceAccount(priceAccount)
}
