// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateUserGovTokenInsuranceStake is the `updateUserGovTokenInsuranceStake` instruction.
type UpdateUserGovTokenInsuranceStake struct {

	// [0] = [] state
	//
	// [1] = [WRITE] spotMarket
	//
	// [2] = [WRITE] insuranceFundStake
	//
	// [3] = [WRITE] userStats
	//
	// [4] = [SIGNER] signer
	//
	// [5] = [WRITE] insuranceFundVault
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateUserGovTokenInsuranceStakeInstructionBuilder creates a new `UpdateUserGovTokenInsuranceStake` instruction builder.
func NewUpdateUserGovTokenInsuranceStakeInstructionBuilder() *UpdateUserGovTokenInsuranceStake {
	nd := &UpdateUserGovTokenInsuranceStake{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *UpdateUserGovTokenInsuranceStake) SetStateAccount(state ag_solanago.PublicKey) *UpdateUserGovTokenInsuranceStake {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateUserGovTokenInsuranceStake) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *UpdateUserGovTokenInsuranceStake) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *UpdateUserGovTokenInsuranceStake {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *UpdateUserGovTokenInsuranceStake) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetInsuranceFundStakeAccount sets the "insuranceFundStake" account.
func (inst *UpdateUserGovTokenInsuranceStake) SetInsuranceFundStakeAccount(insuranceFundStake ag_solanago.PublicKey) *UpdateUserGovTokenInsuranceStake {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(insuranceFundStake).WRITE()
	return inst
}

// GetInsuranceFundStakeAccount gets the "insuranceFundStake" account.
func (inst *UpdateUserGovTokenInsuranceStake) GetInsuranceFundStakeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserStatsAccount sets the "userStats" account.
func (inst *UpdateUserGovTokenInsuranceStake) SetUserStatsAccount(userStats ag_solanago.PublicKey) *UpdateUserGovTokenInsuranceStake {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userStats).WRITE()
	return inst
}

// GetUserStatsAccount gets the "userStats" account.
func (inst *UpdateUserGovTokenInsuranceStake) GetUserStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSignerAccount sets the "signer" account.
func (inst *UpdateUserGovTokenInsuranceStake) SetSignerAccount(signer ag_solanago.PublicKey) *UpdateUserGovTokenInsuranceStake {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *UpdateUserGovTokenInsuranceStake) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetInsuranceFundVaultAccount sets the "insuranceFundVault" account.
func (inst *UpdateUserGovTokenInsuranceStake) SetInsuranceFundVaultAccount(insuranceFundVault ag_solanago.PublicKey) *UpdateUserGovTokenInsuranceStake {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(insuranceFundVault).WRITE()
	return inst
}

// GetInsuranceFundVaultAccount gets the "insuranceFundVault" account.
func (inst *UpdateUserGovTokenInsuranceStake) GetInsuranceFundVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst UpdateUserGovTokenInsuranceStake) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateUserGovTokenInsuranceStake,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateUserGovTokenInsuranceStake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateUserGovTokenInsuranceStake) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SpotMarket is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.InsuranceFundStake is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserStats is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Signer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.InsuranceFundVault is not set")
		}
	}
	return nil
}

func (inst *UpdateUserGovTokenInsuranceStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateUserGovTokenInsuranceStake")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        spotMarket", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("insuranceFundStake", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         userStats", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            signer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("insuranceFundVault", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj UpdateUserGovTokenInsuranceStake) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UpdateUserGovTokenInsuranceStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUpdateUserGovTokenInsuranceStakeInstruction declares a new UpdateUserGovTokenInsuranceStake instruction with the provided parameters and accounts.
func NewUpdateUserGovTokenInsuranceStakeInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	spotMarket ag_solanago.PublicKey,
	insuranceFundStake ag_solanago.PublicKey,
	userStats ag_solanago.PublicKey,
	signer ag_solanago.PublicKey,
	insuranceFundVault ag_solanago.PublicKey) *UpdateUserGovTokenInsuranceStake {
	return NewUpdateUserGovTokenInsuranceStakeInstructionBuilder().
		SetStateAccount(state).
		SetSpotMarketAccount(spotMarket).
		SetInsuranceFundStakeAccount(insuranceFundStake).
		SetUserStatsAccount(userStats).
		SetSignerAccount(signer).
		SetInsuranceFundVaultAccount(insuranceFundVault)
}
