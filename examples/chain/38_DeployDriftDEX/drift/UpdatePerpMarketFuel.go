// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdatePerpMarketFuel is the `updatePerpMarketFuel` instruction.
type UpdatePerpMarketFuel struct {
	FuelBoostTaker    *uint8 `bin:"optional"`
	FuelBoostMaker    *uint8 `bin:"optional"`
	FuelBoostPosition *uint8 `bin:"optional"`

	// [0] = [SIGNER] admin
	//
	// [1] = [] state
	//
	// [2] = [WRITE] perpMarket
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePerpMarketFuelInstructionBuilder creates a new `UpdatePerpMarketFuel` instruction builder.
func NewUpdatePerpMarketFuelInstructionBuilder() *UpdatePerpMarketFuel {
	nd := &UpdatePerpMarketFuel{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetFuelBoostTaker sets the "fuelBoostTaker" parameter.
func (inst *UpdatePerpMarketFuel) SetFuelBoostTaker(fuelBoostTaker uint8) *UpdatePerpMarketFuel {
	inst.FuelBoostTaker = &fuelBoostTaker
	return inst
}

// SetFuelBoostMaker sets the "fuelBoostMaker" parameter.
func (inst *UpdatePerpMarketFuel) SetFuelBoostMaker(fuelBoostMaker uint8) *UpdatePerpMarketFuel {
	inst.FuelBoostMaker = &fuelBoostMaker
	return inst
}

// SetFuelBoostPosition sets the "fuelBoostPosition" parameter.
func (inst *UpdatePerpMarketFuel) SetFuelBoostPosition(fuelBoostPosition uint8) *UpdatePerpMarketFuel {
	inst.FuelBoostPosition = &fuelBoostPosition
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdatePerpMarketFuel) SetAdminAccount(admin ag_solanago.PublicKey) *UpdatePerpMarketFuel {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdatePerpMarketFuel) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *UpdatePerpMarketFuel) SetStateAccount(state ag_solanago.PublicKey) *UpdatePerpMarketFuel {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdatePerpMarketFuel) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPerpMarketAccount sets the "perpMarket" account.
func (inst *UpdatePerpMarketFuel) SetPerpMarketAccount(perpMarket ag_solanago.PublicKey) *UpdatePerpMarketFuel {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(perpMarket).WRITE()
	return inst
}

// GetPerpMarketAccount gets the "perpMarket" account.
func (inst *UpdatePerpMarketFuel) GetPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdatePerpMarketFuel) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePerpMarketFuel,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePerpMarketFuel) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePerpMarketFuel) Validate() error {
	// Check whether all (required) parameters are set:
	{
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

func (inst *UpdatePerpMarketFuel) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePerpMarketFuel")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   FuelBoostTaker (OPT)", inst.FuelBoostTaker))
						paramsBranch.Child(ag_format.Param("   FuelBoostMaker (OPT)", inst.FuelBoostMaker))
						paramsBranch.Child(ag_format.Param("FuelBoostPosition (OPT)", inst.FuelBoostPosition))
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

func (obj UpdatePerpMarketFuel) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `FuelBoostTaker` param (optional):
	{
		if obj.FuelBoostTaker == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.FuelBoostTaker)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FuelBoostMaker` param (optional):
	{
		if obj.FuelBoostMaker == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.FuelBoostMaker)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FuelBoostPosition` param (optional):
	{
		if obj.FuelBoostPosition == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.FuelBoostPosition)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *UpdatePerpMarketFuel) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `FuelBoostTaker` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.FuelBoostTaker)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FuelBoostMaker` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.FuelBoostMaker)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FuelBoostPosition` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.FuelBoostPosition)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewUpdatePerpMarketFuelInstruction declares a new UpdatePerpMarketFuel instruction with the provided parameters and accounts.
func NewUpdatePerpMarketFuelInstruction(
	// Parameters:
	fuelBoostTaker uint8,
	fuelBoostMaker uint8,
	fuelBoostPosition uint8,
	// Accounts:
	admin ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	perpMarket ag_solanago.PublicKey) *UpdatePerpMarketFuel {
	return NewUpdatePerpMarketFuelInstructionBuilder().
		SetFuelBoostTaker(fuelBoostTaker).
		SetFuelBoostMaker(fuelBoostMaker).
		SetFuelBoostPosition(fuelBoostPosition).
		SetAdminAccount(admin).
		SetStateAccount(state).
		SetPerpMarketAccount(perpMarket)
}