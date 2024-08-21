// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdatePerpMarketMaxImbalances is the `updatePerpMarketMaxImbalances` instruction.
type UpdatePerpMarketMaxImbalances struct {
	UnrealizedMaxImbalance      *uint64
	MaxRevenueWithdrawPerPeriod *uint64
	QuoteMaxInsurance           *uint64

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] perpMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePerpMarketMaxImbalancesInstructionBuilder creates a new `UpdatePerpMarketMaxImbalances` instruction builder.
func NewUpdatePerpMarketMaxImbalancesInstructionBuilder() *UpdatePerpMarketMaxImbalances {
	nd := &UpdatePerpMarketMaxImbalances{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetUnrealizedMaxImbalance sets the "unrealizedMaxImbalance" parameter.
func (inst *UpdatePerpMarketMaxImbalances) SetUnrealizedMaxImbalance(unrealizedMaxImbalance uint64) *UpdatePerpMarketMaxImbalances {
	inst.UnrealizedMaxImbalance = &unrealizedMaxImbalance
	return inst
}

// SetMaxRevenueWithdrawPerPeriod sets the "maxRevenueWithdrawPerPeriod" parameter.
func (inst *UpdatePerpMarketMaxImbalances) SetMaxRevenueWithdrawPerPeriod(maxRevenueWithdrawPerPeriod uint64) *UpdatePerpMarketMaxImbalances {
	inst.MaxRevenueWithdrawPerPeriod = &maxRevenueWithdrawPerPeriod
	return inst
}

// SetQuoteMaxInsurance sets the "quoteMaxInsurance" parameter.
func (inst *UpdatePerpMarketMaxImbalances) SetQuoteMaxInsurance(quoteMaxInsurance uint64) *UpdatePerpMarketMaxImbalances {
	inst.QuoteMaxInsurance = &quoteMaxInsurance
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdatePerpMarketMaxImbalances) SetAdminAccount(admin ag_solanago.PublicKey) *UpdatePerpMarketMaxImbalances {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdatePerpMarketMaxImbalances) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdatePerpMarketMaxImbalances) SetStateAccount(state ag_solanago.PublicKey) *UpdatePerpMarketMaxImbalances {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdatePerpMarketMaxImbalances) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *UpdatePerpMarketMaxImbalances) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *UpdatePerpMarketMaxImbalances {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *UpdatePerpMarketMaxImbalances) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdatePerpMarketMaxImbalances) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePerpMarketMaxImbalances,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePerpMarketMaxImbalances) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePerpMarketMaxImbalances) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UnrealizedMaxImbalance == nil {
			return errors.New("UnrealizedMaxImbalance parameter is not set")
		}
		if inst.MaxRevenueWithdrawPerPeriod == nil {
			return errors.New("MaxRevenueWithdrawPerPeriod parameter is not set")
		}
		if inst.QuoteMaxInsurance == nil {
			return errors.New("QuoteMaxInsurance parameter is not set")
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

func (inst *UpdatePerpMarketMaxImbalances) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePerpMarketMaxImbalances")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     UnrealizedMaxImbalance", *inst.UnrealizedMaxImbalance))
						paramsBranch.Child(ag_format.Param("MaxRevenueWithdrawPerPeriod", *inst.MaxRevenueWithdrawPerPeriod))
						paramsBranch.Child(ag_format.Param("          QuoteMaxInsurance", *inst.QuoteMaxInsurance))
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

func (obj UpdatePerpMarketMaxImbalances) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UnrealizedMaxImbalance` param:
	err = encoder.Encode(obj.UnrealizedMaxImbalance)
	if err != nil {
		return err
	}
	// Serialize `MaxRevenueWithdrawPerPeriod` param:
	err = encoder.Encode(obj.MaxRevenueWithdrawPerPeriod)
	if err != nil {
		return err
	}
	// Serialize `QuoteMaxInsurance` param:
	err = encoder.Encode(obj.QuoteMaxInsurance)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdatePerpMarketMaxImbalances) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UnrealizedMaxImbalance`:
	err = decoder.Decode(&obj.UnrealizedMaxImbalance)
	if err != nil {
		return err
	}
	// Deserialize `MaxRevenueWithdrawPerPeriod`:
	err = decoder.Decode(&obj.MaxRevenueWithdrawPerPeriod)
	if err != nil {
		return err
	}
	// Deserialize `QuoteMaxInsurance`:
	err = decoder.Decode(&obj.QuoteMaxInsurance)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdatePerpMarketMaxImbalancesInstruction declares a new UpdatePerpMarketMaxImbalances instruction with the provided parameters and accounts.
func NewUpdatePerpMarketMaxImbalancesInstruction(
	// Parameters:
	unrealizedMaxImbalance uint64,
	maxRevenueWithdrawPerPeriod uint64,
	quoteMaxInsurance uint64,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey) *UpdatePerpMarketMaxImbalances {
	return NewUpdatePerpMarketMaxImbalancesInstructionBuilder().
		SetUnrealizedMaxImbalance(unrealizedMaxImbalance).
		SetMaxRevenueWithdrawPerPeriod(maxRevenueWithdrawPerPeriod).
		SetQuoteMaxInsurance(quoteMaxInsurance).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket)
}