// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ResolveSpotBankruptcy is the `resolveSpotBankruptcy` instruction.
type ResolveSpotBankruptcy struct {
	MarketIndex *uint16

	// [0] = [] state
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [WRITE] liquidator
	//
	// [3] = [WRITE] liquidatorStats
	//
	// [4] = [WRITE] user
	//
	// [5] = [WRITE] userStats
	//
	// [6] = [WRITE] spotMarketVault
	//
	// [7] = [WRITE] insuranceFundVault
	//
	// [8] = [] driftSigner
	//
	// [9] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewResolveSpotBankruptcyInstructionBuilder creates a new `ResolveSpotBankruptcy` instruction builder.
func NewResolveSpotBankruptcyInstructionBuilder() *ResolveSpotBankruptcy {
	nd := &ResolveSpotBankruptcy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetMarketIndex sets the "marketIndex" parameter.
func (inst *ResolveSpotBankruptcy) SetMarketIndex(marketIndex uint16) *ResolveSpotBankruptcy {
	inst.MarketIndex = &marketIndex
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *ResolveSpotBankruptcy) SetStateAccount(state ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *ResolveSpotBankruptcy) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ResolveSpotBankruptcy) SetAuthorityAccount(authority ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ResolveSpotBankruptcy) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLiquidatorAccount sets the "liquidator" account.
func (inst *ResolveSpotBankruptcy) SetLiquidatorAccount(liquidator ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(liquidator).WRITE()
	return inst
}

// GetLiquidatorAccount gets the "liquidator" account.
func (inst *ResolveSpotBankruptcy) GetLiquidatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLiquidatorStatsAccount sets the "liquidatorStats" account.
func (inst *ResolveSpotBankruptcy) SetLiquidatorStatsAccount(liquidatorStats ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(liquidatorStats).WRITE()
	return inst
}

// GetLiquidatorStatsAccount gets the "liquidatorStats" account.
func (inst *ResolveSpotBankruptcy) GetLiquidatorStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserAccount sets the "user" account.
func (inst *ResolveSpotBankruptcy) SetUserAccount(user ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(user).WRITE()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *ResolveSpotBankruptcy) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetUserStatsAccount sets the "userStats" account.
func (inst *ResolveSpotBankruptcy) SetUserStatsAccount(userStats ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(userStats).WRITE()
	return inst
}

// GetUserStatsAccount gets the "userStats" account.
func (inst *ResolveSpotBankruptcy) GetUserStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSpotMarketVaultAccount sets the "spotMarketVault" account.
func (inst *ResolveSpotBankruptcy) SetSpotMarketVaultAccount(spotMarketVault ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(spotMarketVault).WRITE()
	return inst
}

// GetSpotMarketVaultAccount gets the "spotMarketVault" account.
func (inst *ResolveSpotBankruptcy) GetSpotMarketVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetInsuranceFundVaultAccount sets the "insuranceFundVault" account.
func (inst *ResolveSpotBankruptcy) SetInsuranceFundVaultAccount(insuranceFundVault ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(insuranceFundVault).WRITE()
	return inst
}

// GetInsuranceFundVaultAccount gets the "insuranceFundVault" account.
func (inst *ResolveSpotBankruptcy) GetInsuranceFundVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetDriftSignerAccount sets the "driftSigner" account.
func (inst *ResolveSpotBankruptcy) SetDriftSignerAccount(driftSigner ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(driftSigner)
	return inst
}

// GetDriftSignerAccount gets the "driftSigner" account.
func (inst *ResolveSpotBankruptcy) GetDriftSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *ResolveSpotBankruptcy) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *ResolveSpotBankruptcy) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst ResolveSpotBankruptcy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ResolveSpotBankruptcy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ResolveSpotBankruptcy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ResolveSpotBankruptcy) Validate() error {
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
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Liquidator is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LiquidatorStats is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.UserStats is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SpotMarketVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.InsuranceFundVault is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.DriftSigner is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *ResolveSpotBankruptcy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ResolveSpotBankruptcy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketIndex", *inst.MarketIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        liquidator", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   liquidatorStats", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("              user", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         userStats", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("   spotMarketVault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("insuranceFundVault", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       driftSigner", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj ResolveSpotBankruptcy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketIndex` param:
	err = encoder.Encode(obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ResolveSpotBankruptcy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketIndex`:
	err = decoder.Decode(&obj.MarketIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewResolveSpotBankruptcyInstruction declares a new ResolveSpotBankruptcy instruction with the provided parameters and accounts.
func NewResolveSpotBankruptcyInstruction(
	// Parameters:
	marketIndex uint16,
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	liquidator ag_solanago.PublicKey,
	liquidatorStats ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	userStats ag_solanago.PublicKey,
	spotMarketVault ag_solanago.PublicKey,
	insuranceFundVault ag_solanago.PublicKey,
	driftSigner ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *ResolveSpotBankruptcy {
	return NewResolveSpotBankruptcyInstructionBuilder().
		SetMarketIndex(marketIndex).
		SetStateAccount(state).
		SetAuthorityAccount(authority).
		SetLiquidatorAccount(liquidator).
		SetLiquidatorStatsAccount(liquidatorStats).
		SetUserAccount(user).
		SetUserStatsAccount(userStats).
		SetSpotMarketVaultAccount(spotMarketVault).
		SetInsuranceFundVaultAccount(insuranceFundVault).
		SetDriftSignerAccount(driftSigner).
		SetTokenProgramAccount(tokenProgram)
}