// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelOrderByUserId is the `cancelOrderByUserId` instruction.
type CancelOrderByUserId struct {
	UserOrderId *uint8

	// [0] = [] state
	//
	// [1] = [WRITE] user
	//
	// [2] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelOrderByUserIdInstructionBuilder creates a new `CancelOrderByUserId` instruction builder.
func NewCancelOrderByUserIdInstructionBuilder() *CancelOrderByUserId {
	nd := &CancelOrderByUserId{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetUserOrderId sets the "userOrderId" parameter.
func (inst *CancelOrderByUserId) SetUserOrderId(userOrderId uint8) *CancelOrderByUserId {
	inst.UserOrderId = &userOrderId
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *CancelOrderByUserId) SetStateAccount(state ag_solanago.PublicKey) *CancelOrderByUserId {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *CancelOrderByUserId) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserAccount sets the "user" account.
func (inst *CancelOrderByUserId) SetUserAccount(user ag_solanago.PublicKey) *CancelOrderByUserId {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(user).WRITE()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *CancelOrderByUserId) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *CancelOrderByUserId) SetAuthorityAccount(authority ag_solanago.PublicKey) *CancelOrderByUserId {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *CancelOrderByUserId) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst CancelOrderByUserId) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelOrderByUserId,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelOrderByUserId) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelOrderByUserId) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UserOrderId == nil {
			return errors.New("UserOrderId parameter is not set")
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
	}
	return nil
}

func (inst *CancelOrderByUserId) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelOrderByUserId")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UserOrderId", *inst.UserOrderId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     user", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CancelOrderByUserId) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UserOrderId` param:
	err = encoder.Encode(obj.UserOrderId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelOrderByUserId) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UserOrderId`:
	err = decoder.Decode(&obj.UserOrderId)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelOrderByUserIdInstruction declares a new CancelOrderByUserId instruction with the provided parameters and accounts.
func NewCancelOrderByUserIdInstruction(
	// Parameters:
	userOrderId uint8,
	// Accounts:
	state ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *CancelOrderByUserId {
	return NewCancelOrderByUserIdInstructionBuilder().
		SetUserOrderId(userOrderId).
		SetStateAccount(state).
		SetUserAccount(user).
		SetAuthorityAccount(authority)
}
