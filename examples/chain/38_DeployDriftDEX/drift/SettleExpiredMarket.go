// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SettleExpiredMarket is the `settleExpiredMarket` instruction.
type SettleExpiredMarket struct {
	MarketIndex *uint16

	// [0] = [] state
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSettleExpiredMarketInstructionBuilder creates a new `SettleExpiredMarket` instruction builder.
func NewSettleExpiredMarketInstructionBuilder() *SettleExpiredMarket {
	nd := &SettleExpiredMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMarketIndex sets the "marketIndex" parameter.
func (inst *SettleExpiredMarket) SetMarketIndex(marketIndex uint16) *SettleExpiredMarket {
	inst.MarketIndex = &marketIndex
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *SettleExpiredMarket) SetStateAccount(state ag_solanago.PublicKey) *SettleExpiredMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SettleExpiredMarket) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SettleExpiredMarket) SetAuthorityAccount(authority ag_solanago.PublicKey) *SettleExpiredMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SettleExpiredMarket) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst SettleExpiredMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SettleExpiredMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SettleExpiredMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SettleExpiredMarket) Validate() error {
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
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *SettleExpiredMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SettleExpiredMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketIndex", *inst.MarketIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj SettleExpiredMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketIndex` param:
	err = encoder.Encode(obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SettleExpiredMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketIndex`:
	err = decoder.Decode(&obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewSettleExpiredMarketInstruction declares a new SettleExpiredMarket instruction with the provided parameters and accounts.
func NewSettleExpiredMarketInstruction(
	// Parameters:
	marketIndex uint16,
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SettleExpiredMarket {
	return NewSettleExpiredMarketInstructionBuilder().
		SetMarketIndex(marketIndex).
		SetStateAccount(state).
		SetAuthorityAccount(authority)
}
