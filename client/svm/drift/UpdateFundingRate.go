// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateFundingRate is the `updateFundingRate` instruction.
type UpdateFundingRate struct {
	MarketIndex *uint16

	// [0] = [] state
	//
	// [1] = [WRITE] perpMarket
	//
	// [2] = [] oracle
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateFundingRateInstructionBuilder creates a new `UpdateFundingRate` instruction builder.
func NewUpdateFundingRateInstructionBuilder() *UpdateFundingRate {
	nd := &UpdateFundingRate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetMarketIndex sets the "marketIndex" parameter.
func (inst *UpdateFundingRate) SetMarketIndex(marketIndex uint16) *UpdateFundingRate {
	inst.MarketIndex = &marketIndex
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *UpdateFundingRate) SetStateAccount(state ag_solanago.PublicKey) *UpdateFundingRate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateFundingRate) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *UpdateFundingRate) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *UpdateFundingRate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *UpdateFundingRate) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOracleAccount sets the "oracle" account.
func (inst *UpdateFundingRate) SetOracleAccount(oracle ag_solanago.PublicKey) *UpdateFundingRate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(oracle)
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *UpdateFundingRate) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateFundingRate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateFundingRate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateFundingRate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateFundingRate) Validate() error {
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
			return errors.New("accounts.PerpMarket is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Oracle is not set")
		}
	}
	return nil
}

func (inst *UpdateFundingRate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateFundingRate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketIndex", *inst.MarketIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("perpMarket", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    oracle", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdateFundingRate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketIndex` param:
	err = encoder.Encode(obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateFundingRate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketIndex`:
	err = decoder.Decode(&obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateFundingRateInstruction declares a new UpdateFundingRate instruction with the provided parameters and accounts.
func NewUpdateFundingRateInstruction(
	// Parameters:
	marketIndex uint16,
	// Accounts:
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey) *UpdateFundingRate {
	return NewUpdateFundingRateInstructionBuilder().
		SetMarketIndex(marketIndex).
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket).
		SetOracleAccount(oracle)
}