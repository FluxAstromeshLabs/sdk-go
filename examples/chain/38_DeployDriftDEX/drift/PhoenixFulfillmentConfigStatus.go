// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PhoenixFulfillmentConfigStatus is the `phoenixFulfillmentConfigStatus` instruction.
type PhoenixFulfillmentConfigStatus struct {
	Status *SpotFulfillmentConfigStatus

	// [0] = [] state
	//
	// [1] = [WRITE] phoenixFulfillmentConfig
	//
	// [2] = [WRITE, SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPhoenixFulfillmentConfigStatusInstructionBuilder creates a new `PhoenixFulfillmentConfigStatus` instruction builder.
func NewPhoenixFulfillmentConfigStatusInstructionBuilder() *PhoenixFulfillmentConfigStatus {
	nd := &PhoenixFulfillmentConfigStatus{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetStatus sets the "status" parameter.
func (inst *PhoenixFulfillmentConfigStatus) SetStatus(status SpotFulfillmentConfigStatus) *PhoenixFulfillmentConfigStatus {
	inst.Status = &status
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *PhoenixFulfillmentConfigStatus) SetStateAccount(state ag_solanago.PublicKey) *PhoenixFulfillmentConfigStatus {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *PhoenixFulfillmentConfigStatus) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPhoenixFulfillmentConfigAccount sets the "phoenixFulfillmentConfig" account.
func (inst *PhoenixFulfillmentConfigStatus) SetPhoenixFulfillmentConfigAccount(phoenixFulfillmentConfig ag_solanago.PublicKey) *PhoenixFulfillmentConfigStatus {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(phoenixFulfillmentConfig).WRITE()
	return inst
}

// GetPhoenixFulfillmentConfigAccount gets the "phoenixFulfillmentConfig" account.
func (inst *PhoenixFulfillmentConfigStatus) GetPhoenixFulfillmentConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAdminAccount sets the "admin" account.
func (inst *PhoenixFulfillmentConfigStatus) SetAdminAccount(admin ag_solanago.PublicKey) *PhoenixFulfillmentConfigStatus {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *PhoenixFulfillmentConfigStatus) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst PhoenixFulfillmentConfigStatus) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PhoenixFulfillmentConfigStatus,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PhoenixFulfillmentConfigStatus) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PhoenixFulfillmentConfigStatus) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Status == nil {
			return errors.New("Status parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PhoenixFulfillmentConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *PhoenixFulfillmentConfigStatus) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PhoenixFulfillmentConfigStatus")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Status", *inst.Status))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("phoenixFulfillmentConfig", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj PhoenixFulfillmentConfigStatus) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PhoenixFulfillmentConfigStatus) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	return nil
}

// NewPhoenixFulfillmentConfigStatusInstruction declares a new PhoenixFulfillmentConfigStatus instruction with the provided parameters and accounts.
func NewPhoenixFulfillmentConfigStatusInstruction(
	// Parameters:
	status SpotFulfillmentConfigStatus,
	// Accounts:
	state ag_solanago.PublicKey,
	phoenixFulfillmentConfig ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *PhoenixFulfillmentConfigStatus {
	return NewPhoenixFulfillmentConfigStatusInstructionBuilder().
		SetStatus(status).
		SetStateAccount(state).
		SetPhoenixFulfillmentConfigAccount(phoenixFulfillmentConfig).
		SetAdminAccount(admin)
}
