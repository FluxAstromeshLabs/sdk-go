// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelRequestRemoveInsuranceFundStake is the `cancelRequestRemoveInsuranceFundStake` instruction.
type CancelRequestRemoveInsuranceFundStake struct {
	MarketIndex *uint16

	// [0] = [WRITE] spotMarket
	//
	// [1] = [WRITE] insuranceFundStake
	//
	// [2] = [WRITE] userStats
	//
	// [3] = [SIGNER] authority
	//
	// [4] = [WRITE] insuranceFundVault
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelRequestRemoveInsuranceFundStakeInstructionBuilder creates a new `CancelRequestRemoveInsuranceFundStake` instruction builder.
func NewCancelRequestRemoveInsuranceFundStakeInstructionBuilder() *CancelRequestRemoveInsuranceFundStake {
	nd := &CancelRequestRemoveInsuranceFundStake{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetMarketIndex sets the "marketIndex" parameter.
func (inst *CancelRequestRemoveInsuranceFundStake) SetMarketIndex(marketIndex uint16) *CancelRequestRemoveInsuranceFundStake {
	inst.MarketIndex = &marketIndex
	return inst
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *CancelRequestRemoveInsuranceFundStake) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *CancelRequestRemoveInsuranceFundStake {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *CancelRequestRemoveInsuranceFundStake) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetInsuranceFundStakeAccount sets the "insuranceFundStake" account.
func (inst *CancelRequestRemoveInsuranceFundStake) SetInsuranceFundStakeAccount(insuranceFundStake ag_solanago.PublicKey) *CancelRequestRemoveInsuranceFundStake {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(insuranceFundStake).WRITE()
	return inst
}

// GetInsuranceFundStakeAccount gets the "insuranceFundStake" account.
func (inst *CancelRequestRemoveInsuranceFundStake) GetInsuranceFundStakeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserStatsAccount sets the "userStats" account.
func (inst *CancelRequestRemoveInsuranceFundStake) SetUserStatsAccount(userStats ag_solanago.PublicKey) *CancelRequestRemoveInsuranceFundStake {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userStats).WRITE()
	return inst
}

// GetUserStatsAccount gets the "userStats" account.
func (inst *CancelRequestRemoveInsuranceFundStake) GetUserStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *CancelRequestRemoveInsuranceFundStake) SetAuthorityAccount(authority ag_solanago.PublicKey) *CancelRequestRemoveInsuranceFundStake {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *CancelRequestRemoveInsuranceFundStake) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetInsuranceFundVaultAccount sets the "insuranceFundVault" account.
func (inst *CancelRequestRemoveInsuranceFundStake) SetInsuranceFundVaultAccount(insuranceFundVault ag_solanago.PublicKey) *CancelRequestRemoveInsuranceFundStake {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(insuranceFundVault).WRITE()
	return inst
}

// GetInsuranceFundVaultAccount gets the "insuranceFundVault" account.
func (inst *CancelRequestRemoveInsuranceFundStake) GetInsuranceFundVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst CancelRequestRemoveInsuranceFundStake) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelRequestRemoveInsuranceFundStake,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelRequestRemoveInsuranceFundStake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelRequestRemoveInsuranceFundStake) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MarketIndex == nil {
			return errors.New("MarketIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SpotMarket is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.InsuranceFundStake is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserStats is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.InsuranceFundVault is not set")
		}
	}
	return nil
}

func (inst *CancelRequestRemoveInsuranceFundStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelRequestRemoveInsuranceFundStake")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketIndex", *inst.MarketIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        spotMarket", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("insuranceFundStake", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         userStats", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("insuranceFundVault", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj CancelRequestRemoveInsuranceFundStake) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketIndex` param:
	err = encoder.Encode(obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelRequestRemoveInsuranceFundStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketIndex`:
	err = decoder.Decode(&obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelRequestRemoveInsuranceFundStakeInstruction declares a new CancelRequestRemoveInsuranceFundStake instruction with the provided parameters and accounts.
func NewCancelRequestRemoveInsuranceFundStakeInstruction(
	// Parameters:
	marketIndex uint16,
	// Accounts:
	spotMarket ag_solanago.PublicKey,
	insuranceFundStake ag_solanago.PublicKey,
	userStats ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	insuranceFundVault ag_solanago.PublicKey) *CancelRequestRemoveInsuranceFundStake {
	return NewCancelRequestRemoveInsuranceFundStakeInstructionBuilder().
		SetMarketIndex(marketIndex).
		SetSpotMarketAccount(spotMarket).
		SetInsuranceFundStakeAccount(insuranceFundStake).
		SetUserStatsAccount(userStats).
		SetAuthorityAccount(authority).
		SetInsuranceFundVaultAccount(insuranceFundVault)
}