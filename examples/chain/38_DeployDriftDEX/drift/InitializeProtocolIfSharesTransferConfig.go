// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeProtocolIfSharesTransferConfig is the `initializeProtocolIfSharesTransferConfig` instruction.
type InitializeProtocolIfSharesTransferConfig struct {

	// [0] = [WRITE, SIGNER] admin
	//
	// [1] = [WRITE] protocolIfSharesTransferConfig
	//
	// [2] = [] state
	//
	// [3] = [] rent
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeProtocolIfSharesTransferConfigInstructionBuilder creates a new `InitializeProtocolIfSharesTransferConfig` instruction builder.
func NewInitializeProtocolIfSharesTransferConfigInstructionBuilder() *InitializeProtocolIfSharesTransferConfig {
	nd := &InitializeProtocolIfSharesTransferConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetAdminAccount sets the "admin" account.
func (inst *InitializeProtocolIfSharesTransferConfig) SetAdminAccount(admin ag_solanago.PublicKey) *InitializeProtocolIfSharesTransferConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitializeProtocolIfSharesTransferConfig) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetProtocolIfSharesTransferConfigAccount sets the "protocolIfSharesTransferConfig" account.
func (inst *InitializeProtocolIfSharesTransferConfig) SetProtocolIfSharesTransferConfigAccount(protocolIfSharesTransferConfig ag_solanago.PublicKey) *InitializeProtocolIfSharesTransferConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(protocolIfSharesTransferConfig).WRITE()
	return inst
}

// GetProtocolIfSharesTransferConfigAccount gets the "protocolIfSharesTransferConfig" account.
func (inst *InitializeProtocolIfSharesTransferConfig) GetProtocolIfSharesTransferConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStateAccount sets the "state" account.
func (inst *InitializeProtocolIfSharesTransferConfig) SetStateAccount(state ag_solanago.PublicKey) *InitializeProtocolIfSharesTransferConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializeProtocolIfSharesTransferConfig) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeProtocolIfSharesTransferConfig) SetRentAccount(rent ag_solanago.PublicKey) *InitializeProtocolIfSharesTransferConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeProtocolIfSharesTransferConfig) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeProtocolIfSharesTransferConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeProtocolIfSharesTransferConfig {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeProtocolIfSharesTransferConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst InitializeProtocolIfSharesTransferConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeProtocolIfSharesTransferConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeProtocolIfSharesTransferConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeProtocolIfSharesTransferConfig) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ProtocolIfSharesTransferConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitializeProtocolIfSharesTransferConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeProtocolIfSharesTransferConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                         admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("protocolIfSharesTransferConfig", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                         state", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                          rent", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeProtocolIfSharesTransferConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitializeProtocolIfSharesTransferConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeProtocolIfSharesTransferConfigInstruction declares a new InitializeProtocolIfSharesTransferConfig instruction with the provided parameters and accounts.
func NewInitializeProtocolIfSharesTransferConfigInstruction(
	// Accounts:
	admin ag_solanago.PublicKey,
	protocolIfSharesTransferConfig ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeProtocolIfSharesTransferConfig {
	return NewInitializeProtocolIfSharesTransferConfigInstructionBuilder().
		SetAdminAccount(admin).
		SetProtocolIfSharesTransferConfigAccount(protocolIfSharesTransferConfig).
		SetStateAccount(state).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram)
}
