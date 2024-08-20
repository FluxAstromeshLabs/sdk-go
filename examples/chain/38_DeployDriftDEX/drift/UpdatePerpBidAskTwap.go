// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdatePerpBidAskTwap is the `updatePerpBidAskTwap` instruction.
type UpdatePerpBidAskTwap struct {

	// [0] = [] state
	//
	// [1] = [WRITE] perpMarket
	//
	// [2] = [] oracle
	//
	// [3] = [] keeperStats
	//
	// [4] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePerpBidAskTwapInstructionBuilder creates a new `UpdatePerpBidAskTwap` instruction builder.
func NewUpdatePerpBidAskTwapInstructionBuilder() *UpdatePerpBidAskTwap {
	nd := &UpdatePerpBidAskTwap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *UpdatePerpBidAskTwap) SetStateAccount(state ag_solanago.PublicKey) *UpdatePerpBidAskTwap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdatePerpBidAskTwap) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *UpdatePerpBidAskTwap) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *UpdatePerpBidAskTwap {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *UpdatePerpBidAskTwap) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOracleAccount sets the "oracle" account.
func (inst *UpdatePerpBidAskTwap) SetOracleAccount(oracle ag_solanago.PublicKey) *UpdatePerpBidAskTwap {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(oracle)
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *UpdatePerpBidAskTwap) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetKeeperStatsAccount sets the "keeperStats" account.
func (inst *UpdatePerpBidAskTwap) SetKeeperStatsAccount(keeperStats ag_solanago.PublicKey) *UpdatePerpBidAskTwap {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(keeperStats)
	return inst
}

// GetKeeperStatsAccount gets the "keeperStats" account.
func (inst *UpdatePerpBidAskTwap) GetKeeperStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdatePerpBidAskTwap) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdatePerpBidAskTwap {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdatePerpBidAskTwap) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdatePerpBidAskTwap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePerpBidAskTwap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePerpBidAskTwap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePerpBidAskTwap) Validate() error {
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
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.KeeperStats is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *UpdatePerpBidAskTwap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePerpBidAskTwap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" perpMarket", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     oracle", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("keeperStats", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  authority", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UpdatePerpBidAskTwap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UpdatePerpBidAskTwap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUpdatePerpBidAskTwapInstruction declares a new UpdatePerpBidAskTwap instruction with the provided parameters and accounts.
func NewUpdatePerpBidAskTwapInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	keeperStats ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdatePerpBidAskTwap {
	return NewUpdatePerpBidAskTwapInstructionBuilder().
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket).
		SetOracleAccount(oracle).
		SetKeeperStatsAccount(keeperStats).
		SetAuthorityAccount(authority)
}
