// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SettleExpiredMarketPoolsToRevenuePool is the `settleExpiredMarketPoolsToRevenuePool` instruction.
type SettleExpiredMarketPoolsToRevenuePool struct {

	// [0] = [] state
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [WRITE] spotMarket
	//
	// [3] = [WRITE] perpMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSettleExpiredMarketPoolsToRevenuePoolInstructionBuilder creates a new `SettleExpiredMarketPoolsToRevenuePool` instruction builder.
func NewSettleExpiredMarketPoolsToRevenuePoolInstructionBuilder() *SettleExpiredMarketPoolsToRevenuePool {
	nd := &SettleExpiredMarketPoolsToRevenuePool{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) SetStateAccount(state ag_solanago.PublicKey) *SettleExpiredMarketPoolsToRevenuePool {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) SetAdminAccount(admin ag_solanago.PublicKey) *SettleExpiredMarketPoolsToRevenuePool {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *SettleExpiredMarketPoolsToRevenuePool {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *SettleExpiredMarketPoolsToRevenuePool {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *SettleExpiredMarketPoolsToRevenuePool) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst SettleExpiredMarketPoolsToRevenuePool) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SettleExpiredMarketPoolsToRevenuePool,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SettleExpiredMarketPoolsToRevenuePool) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SettleExpiredMarketPoolsToRevenuePool) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SpotMarket is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PerpMarket is not set")
		}
	}
	return nil
}

func (inst *SettleExpiredMarketPoolsToRevenuePool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SettleExpiredMarketPoolsToRevenuePool")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("spotMarket", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("perpMarket", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj SettleExpiredMarketPoolsToRevenuePool) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SettleExpiredMarketPoolsToRevenuePool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSettleExpiredMarketPoolsToRevenuePoolInstruction declares a new SettleExpiredMarketPoolsToRevenuePool instruction with the provided parameters and accounts.
func NewSettleExpiredMarketPoolsToRevenuePoolInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	spotMarket ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey) *SettleExpiredMarketPoolsToRevenuePool {
	return NewSettleExpiredMarketPoolsToRevenuePoolInstructionBuilder().
		SetStateAccount(state).
		SetAdminAccount(admin).
		SetSpotMarketAccount(spotMarket).
		SetPerpMarketAccount(perpMarket)
}
