// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_cp_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// # Arguments
//
// * `ctx`- The accounts needed by instruction.
// * `index` - The index of amm config, there may be multiple config.
// * `trade_fee_rate` - Trade fee rate, can be changed.
// * `protocol_fee_rate` - The rate of protocol fee within tarde fee.
// * `fund_fee_rate` - The rate of fund fee within tarde fee.
//
type CreateAmmConfig struct {
	Index           *uint16
	TradeFeeRate    *uint64
	ProtocolFeeRate *uint64
	FundFeeRate     *uint64
	CreatePoolFee   *uint64

	// [0] = [WRITE, SIGNER] owner
	// ··········· Address to be set as protocol owner.
	//
	// [1] = [WRITE] ammConfig
	// ··········· Initialize config state account to store protocol owner address and fee rates.
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateAmmConfigInstructionBuilder creates a new `CreateAmmConfig` instruction builder.
func NewCreateAmmConfigInstructionBuilder() *CreateAmmConfig {
	nd := &CreateAmmConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetIndex sets the "index" parameter.
func (inst *CreateAmmConfig) SetIndex(index uint16) *CreateAmmConfig {
	inst.Index = &index
	return inst
}

// SetTradeFeeRate sets the "tradeFeeRate" parameter.
func (inst *CreateAmmConfig) SetTradeFeeRate(tradeFeeRate uint64) *CreateAmmConfig {
	inst.TradeFeeRate = &tradeFeeRate
	return inst
}

// SetProtocolFeeRate sets the "protocolFeeRate" parameter.
func (inst *CreateAmmConfig) SetProtocolFeeRate(protocolFeeRate uint64) *CreateAmmConfig {
	inst.ProtocolFeeRate = &protocolFeeRate
	return inst
}

// SetFundFeeRate sets the "fundFeeRate" parameter.
func (inst *CreateAmmConfig) SetFundFeeRate(fundFeeRate uint64) *CreateAmmConfig {
	inst.FundFeeRate = &fundFeeRate
	return inst
}

// SetCreatePoolFee sets the "createPoolFee" parameter.
func (inst *CreateAmmConfig) SetCreatePoolFee(createPoolFee uint64) *CreateAmmConfig {
	inst.CreatePoolFee = &createPoolFee
	return inst
}

// SetOwnerAccount sets the "owner" account.
// Address to be set as protocol owner.
func (inst *CreateAmmConfig) SetOwnerAccount(owner ag_solanago.PublicKey) *CreateAmmConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// Address to be set as protocol owner.
func (inst *CreateAmmConfig) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAmmConfigAccount sets the "ammConfig" account.
// Initialize config state account to store protocol owner address and fee rates.
func (inst *CreateAmmConfig) SetAmmConfigAccount(ammConfig ag_solanago.PublicKey) *CreateAmmConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ammConfig).WRITE()
	return inst
}

// GetAmmConfigAccount gets the "ammConfig" account.
// Initialize config state account to store protocol owner address and fee rates.
func (inst *CreateAmmConfig) GetAmmConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateAmmConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateAmmConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateAmmConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst CreateAmmConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateAmmConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateAmmConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateAmmConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Index == nil {
			return errors.New("Index parameter is not set")
		}
		if inst.TradeFeeRate == nil {
			return errors.New("TradeFeeRate parameter is not set")
		}
		if inst.ProtocolFeeRate == nil {
			return errors.New("ProtocolFeeRate parameter is not set")
		}
		if inst.FundFeeRate == nil {
			return errors.New("FundFeeRate parameter is not set")
		}
		if inst.CreatePoolFee == nil {
			return errors.New("CreatePoolFee parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AmmConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CreateAmmConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateAmmConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("          Index", *inst.Index))
						paramsBranch.Child(ag_format.Param("   TradeFeeRate", *inst.TradeFeeRate))
						paramsBranch.Child(ag_format.Param("ProtocolFeeRate", *inst.ProtocolFeeRate))
						paramsBranch.Child(ag_format.Param("    FundFeeRate", *inst.FundFeeRate))
						paramsBranch.Child(ag_format.Param("  CreatePoolFee", *inst.CreatePoolFee))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    ammConfig", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CreateAmmConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `TradeFeeRate` param:
	err = encoder.Encode(obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeeRate` param:
	err = encoder.Encode(obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Serialize `FundFeeRate` param:
	err = encoder.Encode(obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Serialize `CreatePoolFee` param:
	err = encoder.Encode(obj.CreatePoolFee)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateAmmConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `TradeFeeRate`:
	err = decoder.Decode(&obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeeRate`:
	err = decoder.Decode(&obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `FundFeeRate`:
	err = decoder.Decode(&obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `CreatePoolFee`:
	err = decoder.Decode(&obj.CreatePoolFee)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateAmmConfigInstruction declares a new CreateAmmConfig instruction with the provided parameters and accounts.
func NewCreateAmmConfigInstruction(
	// Parameters:
	index uint16,
	tradeFeeRate uint64,
	protocolFeeRate uint64,
	fundFeeRate uint64,
	createPoolFee uint64,
	// Accounts:
	owner ag_solanago.PublicKey,
	ammConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateAmmConfig {
	return NewCreateAmmConfigInstructionBuilder().
		SetIndex(index).
		SetTradeFeeRate(tradeFeeRate).
		SetProtocolFeeRate(protocolFeeRate).
		SetFundFeeRate(fundFeeRate).
		SetCreatePoolFee(createPoolFee).
		SetOwnerAccount(owner).
		SetAmmConfigAccount(ammConfig).
		SetSystemProgramAccount(systemProgram)
}
