// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SettleMultiplePnls is the `settleMultiplePnls` instruction.
type SettleMultiplePnls struct {
	MarketIndexes *[]uint16
	Mode          *SettlePnlMode

	// [0] = [] state
	//
	// [1] = [WRITE] user
	//
	// [2] = [SIGNER] authority
	//
	// [3] = [] spotMarketVault
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSettleMultiplePnlsInstructionBuilder creates a new `SettleMultiplePnls` instruction builder.
func NewSettleMultiplePnlsInstructionBuilder() *SettleMultiplePnls {
	nd := &SettleMultiplePnls{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetMarketIndexes sets the "marketIndexes" parameter.
func (inst *SettleMultiplePnls) SetMarketIndexes(marketIndexes []uint16) *SettleMultiplePnls {
	inst.MarketIndexes = &marketIndexes
	return inst
}

// SetMode sets the "mode" parameter.
func (inst *SettleMultiplePnls) SetMode(mode SettlePnlMode) *SettleMultiplePnls {
	inst.Mode = &mode
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *SettleMultiplePnls) SetStateAccount(state ag_solanago.PublicKey) *SettleMultiplePnls {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SettleMultiplePnls) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserAccount sets the "user" account.
func (inst *SettleMultiplePnls) SetUserAccount(user ag_solanago.PublicKey) *SettleMultiplePnls {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(user).WRITE()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *SettleMultiplePnls) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SettleMultiplePnls) SetAuthorityAccount(authority ag_solanago.PublicKey) *SettleMultiplePnls {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SettleMultiplePnls) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSpotMarketVaultAccount sets the "spotMarketVault" account.
func (inst *SettleMultiplePnls) SetSpotMarketVaultAccount(spotMarketVault ag_solanago.PublicKey) *SettleMultiplePnls {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(spotMarketVault)
	return inst
}

// GetSpotMarketVaultAccount gets the "spotMarketVault" account.
func (inst *SettleMultiplePnls) GetSpotMarketVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst SettleMultiplePnls) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SettleMultiplePnls,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SettleMultiplePnls) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SettleMultiplePnls) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MarketIndexes == nil {
			return errors.New("MarketIndexes parameter is not set")
		}
		if inst.Mode == nil {
			return errors.New("Mode parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SpotMarketVault is not set")
		}
	}
	return nil
}

func (inst *SettleMultiplePnls) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SettleMultiplePnls")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketIndexes", *inst.MarketIndexes))
						paramsBranch.Child(ag_format.Param("         Mode", *inst.Mode))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           user", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("spotMarketVault", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj SettleMultiplePnls) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketIndexes` param:
	err = encoder.Encode(obj.MarketIndexes)
	if err != nil {
		return err
	}
	// Serialize `Mode` param:
	err = encoder.Encode(obj.Mode)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SettleMultiplePnls) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketIndexes`:
	err = decoder.Decode(&obj.MarketIndexes)
	if err != nil {
		return err
	}
	// Deserialize `Mode`:
	err = decoder.Decode(&obj.Mode)
	if err != nil {
		return err
	}
	return nil
}

// NewSettleMultiplePnlsInstruction declares a new SettleMultiplePnls instruction with the provided parameters and accounts.
func NewSettleMultiplePnlsInstruction(
	// Parameters:
	marketIndexes []uint16,
	mode SettlePnlMode,
	// Accounts:
	state ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	spotMarketVault ag_solanago.PublicKey) *SettleMultiplePnls {
	return NewSettleMultiplePnlsInstructionBuilder().
		SetMarketIndexes(marketIndexes).
		SetMode(mode).
		SetStateAccount(state).
		SetUserAccount(user).
		SetAuthorityAccount(authority).
		SetSpotMarketVaultAccount(spotMarketVault)
}
