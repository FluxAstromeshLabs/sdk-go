// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// LiquidateBorrowForPerpPnl is the `liquidateBorrowForPerpPnl` instruction.
type LiquidateBorrowForPerpPnl struct {
	PerpMarketIndex                *uint16
	SpotMarketIndex                *uint16
	LiquidatorMaxLiabilityTransfer *ag_binary.Uint128
	LimitPrice                     *uint64 `bin:"optional"`

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
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewLiquidateBorrowForPerpPnlInstructionBuilder creates a new `LiquidateBorrowForPerpPnl` instruction builder.
func NewLiquidateBorrowForPerpPnlInstructionBuilder() *LiquidateBorrowForPerpPnl {
	nd := &LiquidateBorrowForPerpPnl{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetPerpMarketIndex sets the "perpMarketIndex" parameter.
func (inst *LiquidateBorrowForPerpPnl) SetPerpMarketIndex(perpMarketIndex uint16) *LiquidateBorrowForPerpPnl {
	inst.PerpMarketIndex = &perpMarketIndex
	return inst
}

// SetSpotMarketIndex sets the "spotMarketIndex" parameter.
func (inst *LiquidateBorrowForPerpPnl) SetSpotMarketIndex(spotMarketIndex uint16) *LiquidateBorrowForPerpPnl {
	inst.SpotMarketIndex = &spotMarketIndex
	return inst
}

// SetLiquidatorMaxLiabilityTransfer sets the "liquidatorMaxLiabilityTransfer" parameter.
func (inst *LiquidateBorrowForPerpPnl) SetLiquidatorMaxLiabilityTransfer(liquidatorMaxLiabilityTransfer ag_binary.Uint128) *LiquidateBorrowForPerpPnl {
	inst.LiquidatorMaxLiabilityTransfer = &liquidatorMaxLiabilityTransfer
	return inst
}

// SetLimitPrice sets the "limitPrice" parameter.
func (inst *LiquidateBorrowForPerpPnl) SetLimitPrice(limitPrice uint64) *LiquidateBorrowForPerpPnl {
	inst.LimitPrice = &limitPrice
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *LiquidateBorrowForPerpPnl) SetStateAccount(state ag_solanago.PublicKey) *LiquidateBorrowForPerpPnl {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *LiquidateBorrowForPerpPnl) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *LiquidateBorrowForPerpPnl) SetAuthorityAccount(authority ag_solanago.PublicKey) *LiquidateBorrowForPerpPnl {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *LiquidateBorrowForPerpPnl) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLiquidatorAccount sets the "liquidator" account.
func (inst *LiquidateBorrowForPerpPnl) SetLiquidatorAccount(liquidator ag_solanago.PublicKey) *LiquidateBorrowForPerpPnl {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(liquidator).WRITE()
	return inst
}

// GetLiquidatorAccount gets the "liquidator" account.
func (inst *LiquidateBorrowForPerpPnl) GetLiquidatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLiquidatorStatsAccount sets the "liquidatorStats" account.
func (inst *LiquidateBorrowForPerpPnl) SetLiquidatorStatsAccount(liquidatorStats ag_solanago.PublicKey) *LiquidateBorrowForPerpPnl {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(liquidatorStats).WRITE()
	return inst
}

// GetLiquidatorStatsAccount gets the "liquidatorStats" account.
func (inst *LiquidateBorrowForPerpPnl) GetLiquidatorStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserAccount sets the "user" account.
func (inst *LiquidateBorrowForPerpPnl) SetUserAccount(user ag_solanago.PublicKey) *LiquidateBorrowForPerpPnl {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(user).WRITE()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *LiquidateBorrowForPerpPnl) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetUserStatsAccount sets the "userStats" account.
func (inst *LiquidateBorrowForPerpPnl) SetUserStatsAccount(userStats ag_solanago.PublicKey) *LiquidateBorrowForPerpPnl {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(userStats).WRITE()
	return inst
}

// GetUserStatsAccount gets the "userStats" account.
func (inst *LiquidateBorrowForPerpPnl) GetUserStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst LiquidateBorrowForPerpPnl) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_LiquidateBorrowForPerpPnl,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst LiquidateBorrowForPerpPnl) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *LiquidateBorrowForPerpPnl) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PerpMarketIndex == nil {
			return errors.New("PerpMarketIndex parameter is not set")
		}
		if inst.SpotMarketIndex == nil {
			return errors.New("SpotMarketIndex parameter is not set")
		}
		if inst.LiquidatorMaxLiabilityTransfer == nil {
			return errors.New("LiquidatorMaxLiabilityTransfer parameter is not set")
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
	}
	return nil
}

func (inst *LiquidateBorrowForPerpPnl) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("LiquidateBorrowForPerpPnl")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("               PerpMarketIndex", *inst.PerpMarketIndex))
						paramsBranch.Child(ag_format.Param("               SpotMarketIndex", *inst.SpotMarketIndex))
						paramsBranch.Child(ag_format.Param("LiquidatorMaxLiabilityTransfer", *inst.LiquidatorMaxLiabilityTransfer))
						paramsBranch.Child(ag_format.Param("                    LimitPrice (OPT)", inst.LimitPrice))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     liquidator", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("liquidatorStats", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           user", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      userStats", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj LiquidateBorrowForPerpPnl) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PerpMarketIndex` param:
	err = encoder.Encode(obj.PerpMarketIndex)
	if err != nil {
		return err
	}
	// Serialize `SpotMarketIndex` param:
	err = encoder.Encode(obj.SpotMarketIndex)
	if err != nil {
		return err
	}
	// Serialize `LiquidatorMaxLiabilityTransfer` param:
	err = encoder.Encode(obj.LiquidatorMaxLiabilityTransfer)
	if err != nil {
		return err
	}
	// Serialize `LimitPrice` param (optional):
	{
		if obj.LimitPrice == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.LimitPrice)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *LiquidateBorrowForPerpPnl) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PerpMarketIndex`:
	err = decoder.Decode(&obj.PerpMarketIndex)
	if err != nil {
		return err
	}
	// Deserialize `SpotMarketIndex`:
	err = decoder.Decode(&obj.SpotMarketIndex)
	if err != nil {
		return err
	}
	// Deserialize `LiquidatorMaxLiabilityTransfer`:
	err = decoder.Decode(&obj.LiquidatorMaxLiabilityTransfer)
	if err != nil {
		return err
	}
	// Deserialize `LimitPrice` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.LimitPrice)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewLiquidateBorrowForPerpPnlInstruction declares a new LiquidateBorrowForPerpPnl instruction with the provided parameters and accounts.
func NewLiquidateBorrowForPerpPnlInstruction(
	// Parameters:
	perpMarketIndex uint16,
	spotMarketIndex uint16,
	liquidatorMaxLiabilityTransfer ag_binary.Uint128,
	limitPrice uint64,
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	liquidator ag_solanago.PublicKey,
	liquidatorStats ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	userStats ag_solanago.PublicKey) *LiquidateBorrowForPerpPnl {
	return NewLiquidateBorrowForPerpPnlInstructionBuilder().
		SetPerpMarketIndex(perpMarketIndex).
		SetSpotMarketIndex(spotMarketIndex).
		SetLiquidatorMaxLiabilityTransfer(liquidatorMaxLiabilityTransfer).
		SetLimitPrice(limitPrice).
		SetStateAccount(state).
		SetAuthorityAccount(authority).
		SetLiquidatorAccount(liquidator).
		SetLiquidatorStatsAccount(liquidatorStats).
		SetUserAccount(user).
		SetUserStatsAccount(userStats)
}
