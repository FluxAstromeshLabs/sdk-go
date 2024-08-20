// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drift

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeSpotMarket is the `initializeSpotMarket` instruction.
type InitializeSpotMarket struct {
	OptimalUtilization           *uint32
	OptimalBorrowRate            *uint32
	MaxBorrowRate                *uint32
	OracleSource                 *OracleSource
	InitialAssetWeight           *uint32
	MaintenanceAssetWeight       *uint32
	InitialLiabilityWeight       *uint32
	MaintenanceLiabilityWeight   *uint32
	ImfFactor                    *uint32
	LiquidatorFee                *uint32
	IfLiquidationFee             *uint32
	ActiveStatus                 *bool
	AssetTier                    *AssetTier
	ScaleInitialAssetWeightStart *uint64
	WithdrawGuardThreshold       *uint64
	OrderTickSize                *uint64
	OrderStepSize                *uint64
	IfTotalFactor                *uint32
	Name                         *[32]uint8

	// [0] = [WRITE] spotMarket
	//
	// [1] = [] spotMarketMint
	//
	// [2] = [WRITE] spotMarketVault
	//
	// [3] = [WRITE] insuranceFundVault
	//
	// [4] = [] driftSigner
	//
	// [5] = [WRITE] state
	//
	// [6] = [] oracle
	//
	// [7] = [WRITE, SIGNER] admin
	//
	// [8] = [] rent
	//
	// [9] = [] systemProgram
	//
	// [10] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeSpotMarketInstructionBuilder creates a new `InitializeSpotMarket` instruction builder.
func NewInitializeSpotMarketInstructionBuilder() *InitializeSpotMarket {
	nd := &InitializeSpotMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetOptimalUtilization sets the "optimalUtilization" parameter.
func (inst *InitializeSpotMarket) SetOptimalUtilization(optimalUtilization uint32) *InitializeSpotMarket {
	inst.OptimalUtilization = &optimalUtilization
	return inst
}

// SetOptimalBorrowRate sets the "optimalBorrowRate" parameter.
func (inst *InitializeSpotMarket) SetOptimalBorrowRate(optimalBorrowRate uint32) *InitializeSpotMarket {
	inst.OptimalBorrowRate = &optimalBorrowRate
	return inst
}

// SetMaxBorrowRate sets the "maxBorrowRate" parameter.
func (inst *InitializeSpotMarket) SetMaxBorrowRate(maxBorrowRate uint32) *InitializeSpotMarket {
	inst.MaxBorrowRate = &maxBorrowRate
	return inst
}

// SetOracleSource sets the "oracleSource" parameter.
func (inst *InitializeSpotMarket) SetOracleSource(oracleSource OracleSource) *InitializeSpotMarket {
	inst.OracleSource = &oracleSource
	return inst
}

// SetInitialAssetWeight sets the "initialAssetWeight" parameter.
func (inst *InitializeSpotMarket) SetInitialAssetWeight(initialAssetWeight uint32) *InitializeSpotMarket {
	inst.InitialAssetWeight = &initialAssetWeight
	return inst
}

// SetMaintenanceAssetWeight sets the "maintenanceAssetWeight" parameter.
func (inst *InitializeSpotMarket) SetMaintenanceAssetWeight(maintenanceAssetWeight uint32) *InitializeSpotMarket {
	inst.MaintenanceAssetWeight = &maintenanceAssetWeight
	return inst
}

// SetInitialLiabilityWeight sets the "initialLiabilityWeight" parameter.
func (inst *InitializeSpotMarket) SetInitialLiabilityWeight(initialLiabilityWeight uint32) *InitializeSpotMarket {
	inst.InitialLiabilityWeight = &initialLiabilityWeight
	return inst
}

// SetMaintenanceLiabilityWeight sets the "maintenanceLiabilityWeight" parameter.
func (inst *InitializeSpotMarket) SetMaintenanceLiabilityWeight(maintenanceLiabilityWeight uint32) *InitializeSpotMarket {
	inst.MaintenanceLiabilityWeight = &maintenanceLiabilityWeight
	return inst
}

// SetImfFactor sets the "imfFactor" parameter.
func (inst *InitializeSpotMarket) SetImfFactor(imfFactor uint32) *InitializeSpotMarket {
	inst.ImfFactor = &imfFactor
	return inst
}

// SetLiquidatorFee sets the "liquidatorFee" parameter.
func (inst *InitializeSpotMarket) SetLiquidatorFee(liquidatorFee uint32) *InitializeSpotMarket {
	inst.LiquidatorFee = &liquidatorFee
	return inst
}

// SetIfLiquidationFee sets the "ifLiquidationFee" parameter.
func (inst *InitializeSpotMarket) SetIfLiquidationFee(ifLiquidationFee uint32) *InitializeSpotMarket {
	inst.IfLiquidationFee = &ifLiquidationFee
	return inst
}

// SetActiveStatus sets the "activeStatus" parameter.
func (inst *InitializeSpotMarket) SetActiveStatus(activeStatus bool) *InitializeSpotMarket {
	inst.ActiveStatus = &activeStatus
	return inst
}

// SetAssetTier sets the "assetTier" parameter.
func (inst *InitializeSpotMarket) SetAssetTier(assetTier AssetTier) *InitializeSpotMarket {
	inst.AssetTier = &assetTier
	return inst
}

// SetScaleInitialAssetWeightStart sets the "scaleInitialAssetWeightStart" parameter.
func (inst *InitializeSpotMarket) SetScaleInitialAssetWeightStart(scaleInitialAssetWeightStart uint64) *InitializeSpotMarket {
	inst.ScaleInitialAssetWeightStart = &scaleInitialAssetWeightStart
	return inst
}

// SetWithdrawGuardThreshold sets the "withdrawGuardThreshold" parameter.
func (inst *InitializeSpotMarket) SetWithdrawGuardThreshold(withdrawGuardThreshold uint64) *InitializeSpotMarket {
	inst.WithdrawGuardThreshold = &withdrawGuardThreshold
	return inst
}

// SetOrderTickSize sets the "orderTickSize" parameter.
func (inst *InitializeSpotMarket) SetOrderTickSize(orderTickSize uint64) *InitializeSpotMarket {
	inst.OrderTickSize = &orderTickSize
	return inst
}

// SetOrderStepSize sets the "orderStepSize" parameter.
func (inst *InitializeSpotMarket) SetOrderStepSize(orderStepSize uint64) *InitializeSpotMarket {
	inst.OrderStepSize = &orderStepSize
	return inst
}

// SetIfTotalFactor sets the "ifTotalFactor" parameter.
func (inst *InitializeSpotMarket) SetIfTotalFactor(ifTotalFactor uint32) *InitializeSpotMarket {
	inst.IfTotalFactor = &ifTotalFactor
	return inst
}

// SetName sets the "name" parameter.
func (inst *InitializeSpotMarket) SetName(name [32]uint8) *InitializeSpotMarket {
	inst.Name = &name
	return inst
}

// SetSpotMarketAccount sets the "spotMarket" account.
func (inst *InitializeSpotMarket) SetSpotMarketAccount(spotMarket ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(spotMarket).WRITE()
	return inst
}

// GetSpotMarketAccount gets the "spotMarket" account.
func (inst *InitializeSpotMarket) GetSpotMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSpotMarketMintAccount sets the "spotMarketMint" account.
func (inst *InitializeSpotMarket) SetSpotMarketMintAccount(spotMarketMint ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(spotMarketMint)
	return inst
}

// GetSpotMarketMintAccount gets the "spotMarketMint" account.
func (inst *InitializeSpotMarket) GetSpotMarketMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSpotMarketVaultAccount sets the "spotMarketVault" account.
func (inst *InitializeSpotMarket) SetSpotMarketVaultAccount(spotMarketVault ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(spotMarketVault).WRITE()
	return inst
}

// GetSpotMarketVaultAccount gets the "spotMarketVault" account.
func (inst *InitializeSpotMarket) GetSpotMarketVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetInsuranceFundVaultAccount sets the "insuranceFundVault" account.
func (inst *InitializeSpotMarket) SetInsuranceFundVaultAccount(insuranceFundVault ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(insuranceFundVault).WRITE()
	return inst
}

// GetInsuranceFundVaultAccount gets the "insuranceFundVault" account.
func (inst *InitializeSpotMarket) GetInsuranceFundVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetDriftSignerAccount sets the "driftSigner" account.
func (inst *InitializeSpotMarket) SetDriftSignerAccount(driftSigner ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(driftSigner)
	return inst
}

// GetDriftSignerAccount gets the "driftSigner" account.
func (inst *InitializeSpotMarket) GetDriftSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetStateAccount sets the "state" account.
func (inst *InitializeSpotMarket) SetStateAccount(state ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializeSpotMarket) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetOracleAccount sets the "oracle" account.
func (inst *InitializeSpotMarket) SetOracleAccount(oracle ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(oracle)
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *InitializeSpotMarket) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAdminAccount sets the "admin" account.
func (inst *InitializeSpotMarket) SetAdminAccount(admin ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitializeSpotMarket) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeSpotMarket) SetRentAccount(rent ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeSpotMarket) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeSpotMarket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeSpotMarket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *InitializeSpotMarket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InitializeSpotMarket {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *InitializeSpotMarket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst InitializeSpotMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeSpotMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeSpotMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeSpotMarket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.OptimalUtilization == nil {
			return errors.New("OptimalUtilization parameter is not set")
		}
		if inst.OptimalBorrowRate == nil {
			return errors.New("OptimalBorrowRate parameter is not set")
		}
		if inst.MaxBorrowRate == nil {
			return errors.New("MaxBorrowRate parameter is not set")
		}
		if inst.OracleSource == nil {
			return errors.New("OracleSource parameter is not set")
		}
		if inst.InitialAssetWeight == nil {
			return errors.New("InitialAssetWeight parameter is not set")
		}
		if inst.MaintenanceAssetWeight == nil {
			return errors.New("MaintenanceAssetWeight parameter is not set")
		}
		if inst.InitialLiabilityWeight == nil {
			return errors.New("InitialLiabilityWeight parameter is not set")
		}
		if inst.MaintenanceLiabilityWeight == nil {
			return errors.New("MaintenanceLiabilityWeight parameter is not set")
		}
		if inst.ImfFactor == nil {
			return errors.New("ImfFactor parameter is not set")
		}
		if inst.LiquidatorFee == nil {
			return errors.New("LiquidatorFee parameter is not set")
		}
		if inst.IfLiquidationFee == nil {
			return errors.New("IfLiquidationFee parameter is not set")
		}
		if inst.ActiveStatus == nil {
			return errors.New("ActiveStatus parameter is not set")
		}
		if inst.AssetTier == nil {
			return errors.New("AssetTier parameter is not set")
		}
		if inst.ScaleInitialAssetWeightStart == nil {
			return errors.New("ScaleInitialAssetWeightStart parameter is not set")
		}
		if inst.WithdrawGuardThreshold == nil {
			return errors.New("WithdrawGuardThreshold parameter is not set")
		}
		if inst.OrderTickSize == nil {
			return errors.New("OrderTickSize parameter is not set")
		}
		if inst.OrderStepSize == nil {
			return errors.New("OrderStepSize parameter is not set")
		}
		if inst.IfTotalFactor == nil {
			return errors.New("IfTotalFactor parameter is not set")
		}
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SpotMarket is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SpotMarketMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SpotMarketVault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.InsuranceFundVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.DriftSigner is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *InitializeSpotMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeSpotMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=19]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("          OptimalUtilization", *inst.OptimalUtilization))
						paramsBranch.Child(ag_format.Param("           OptimalBorrowRate", *inst.OptimalBorrowRate))
						paramsBranch.Child(ag_format.Param("               MaxBorrowRate", *inst.MaxBorrowRate))
						paramsBranch.Child(ag_format.Param("                OracleSource", *inst.OracleSource))
						paramsBranch.Child(ag_format.Param("          InitialAssetWeight", *inst.InitialAssetWeight))
						paramsBranch.Child(ag_format.Param("      MaintenanceAssetWeight", *inst.MaintenanceAssetWeight))
						paramsBranch.Child(ag_format.Param("      InitialLiabilityWeight", *inst.InitialLiabilityWeight))
						paramsBranch.Child(ag_format.Param("  MaintenanceLiabilityWeight", *inst.MaintenanceLiabilityWeight))
						paramsBranch.Child(ag_format.Param("                   ImfFactor", *inst.ImfFactor))
						paramsBranch.Child(ag_format.Param("               LiquidatorFee", *inst.LiquidatorFee))
						paramsBranch.Child(ag_format.Param("            IfLiquidationFee", *inst.IfLiquidationFee))
						paramsBranch.Child(ag_format.Param("                ActiveStatus", *inst.ActiveStatus))
						paramsBranch.Child(ag_format.Param("                   AssetTier", *inst.AssetTier))
						paramsBranch.Child(ag_format.Param("ScaleInitialAssetWeightStart", *inst.ScaleInitialAssetWeightStart))
						paramsBranch.Child(ag_format.Param("      WithdrawGuardThreshold", *inst.WithdrawGuardThreshold))
						paramsBranch.Child(ag_format.Param("               OrderTickSize", *inst.OrderTickSize))
						paramsBranch.Child(ag_format.Param("               OrderStepSize", *inst.OrderStepSize))
						paramsBranch.Child(ag_format.Param("               IfTotalFactor", *inst.IfTotalFactor))
						paramsBranch.Child(ag_format.Param("                        Name", *inst.Name))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        spotMarket", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    spotMarketMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   spotMarketVault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("insuranceFundVault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       driftSigner", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             state", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            oracle", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("             admin", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              rent", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj InitializeSpotMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OptimalUtilization` param:
	err = encoder.Encode(obj.OptimalUtilization)
	if err != nil {
		return err
	}
	// Serialize `OptimalBorrowRate` param:
	err = encoder.Encode(obj.OptimalBorrowRate)
	if err != nil {
		return err
	}
	// Serialize `MaxBorrowRate` param:
	err = encoder.Encode(obj.MaxBorrowRate)
	if err != nil {
		return err
	}
	// Serialize `OracleSource` param:
	err = encoder.Encode(obj.OracleSource)
	if err != nil {
		return err
	}
	// Serialize `InitialAssetWeight` param:
	err = encoder.Encode(obj.InitialAssetWeight)
	if err != nil {
		return err
	}
	// Serialize `MaintenanceAssetWeight` param:
	err = encoder.Encode(obj.MaintenanceAssetWeight)
	if err != nil {
		return err
	}
	// Serialize `InitialLiabilityWeight` param:
	err = encoder.Encode(obj.InitialLiabilityWeight)
	if err != nil {
		return err
	}
	// Serialize `MaintenanceLiabilityWeight` param:
	err = encoder.Encode(obj.MaintenanceLiabilityWeight)
	if err != nil {
		return err
	}
	// Serialize `ImfFactor` param:
	err = encoder.Encode(obj.ImfFactor)
	if err != nil {
		return err
	}
	// Serialize `LiquidatorFee` param:
	err = encoder.Encode(obj.LiquidatorFee)
	if err != nil {
		return err
	}
	// Serialize `IfLiquidationFee` param:
	err = encoder.Encode(obj.IfLiquidationFee)
	if err != nil {
		return err
	}
	// Serialize `ActiveStatus` param:
	err = encoder.Encode(obj.ActiveStatus)
	if err != nil {
		return err
	}
	// Serialize `AssetTier` param:
	err = encoder.Encode(obj.AssetTier)
	if err != nil {
		return err
	}
	// Serialize `ScaleInitialAssetWeightStart` param:
	err = encoder.Encode(obj.ScaleInitialAssetWeightStart)
	if err != nil {
		return err
	}
	// Serialize `WithdrawGuardThreshold` param:
	err = encoder.Encode(obj.WithdrawGuardThreshold)
	if err != nil {
		return err
	}
	// Serialize `OrderTickSize` param:
	err = encoder.Encode(obj.OrderTickSize)
	if err != nil {
		return err
	}
	// Serialize `OrderStepSize` param:
	err = encoder.Encode(obj.OrderStepSize)
	if err != nil {
		return err
	}
	// Serialize `IfTotalFactor` param:
	err = encoder.Encode(obj.IfTotalFactor)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeSpotMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OptimalUtilization`:
	err = decoder.Decode(&obj.OptimalUtilization)
	if err != nil {
		return err
	}
	// Deserialize `OptimalBorrowRate`:
	err = decoder.Decode(&obj.OptimalBorrowRate)
	if err != nil {
		return err
	}
	// Deserialize `MaxBorrowRate`:
	err = decoder.Decode(&obj.MaxBorrowRate)
	if err != nil {
		return err
	}
	// Deserialize `OracleSource`:
	err = decoder.Decode(&obj.OracleSource)
	if err != nil {
		return err
	}
	// Deserialize `InitialAssetWeight`:
	err = decoder.Decode(&obj.InitialAssetWeight)
	if err != nil {
		return err
	}
	// Deserialize `MaintenanceAssetWeight`:
	err = decoder.Decode(&obj.MaintenanceAssetWeight)
	if err != nil {
		return err
	}
	// Deserialize `InitialLiabilityWeight`:
	err = decoder.Decode(&obj.InitialLiabilityWeight)
	if err != nil {
		return err
	}
	// Deserialize `MaintenanceLiabilityWeight`:
	err = decoder.Decode(&obj.MaintenanceLiabilityWeight)
	if err != nil {
		return err
	}
	// Deserialize `ImfFactor`:
	err = decoder.Decode(&obj.ImfFactor)
	if err != nil {
		return err
	}
	// Deserialize `LiquidatorFee`:
	err = decoder.Decode(&obj.LiquidatorFee)
	if err != nil {
		return err
	}
	// Deserialize `IfLiquidationFee`:
	err = decoder.Decode(&obj.IfLiquidationFee)
	if err != nil {
		return err
	}
	// Deserialize `ActiveStatus`:
	err = decoder.Decode(&obj.ActiveStatus)
	if err != nil {
		return err
	}
	// Deserialize `AssetTier`:
	err = decoder.Decode(&obj.AssetTier)
	if err != nil {
		return err
	}
	// Deserialize `ScaleInitialAssetWeightStart`:
	err = decoder.Decode(&obj.ScaleInitialAssetWeightStart)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawGuardThreshold`:
	err = decoder.Decode(&obj.WithdrawGuardThreshold)
	if err != nil {
		return err
	}
	// Deserialize `OrderTickSize`:
	err = decoder.Decode(&obj.OrderTickSize)
	if err != nil {
		return err
	}
	// Deserialize `OrderStepSize`:
	err = decoder.Decode(&obj.OrderStepSize)
	if err != nil {
		return err
	}
	// Deserialize `IfTotalFactor`:
	err = decoder.Decode(&obj.IfTotalFactor)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeSpotMarketInstruction declares a new InitializeSpotMarket instruction with the provided parameters and accounts.
func NewInitializeSpotMarketInstruction(
	// Parameters:
	optimalUtilization uint32,
	optimalBorrowRate uint32,
	maxBorrowRate uint32,
	oracleSource OracleSource,
	initialAssetWeight uint32,
	maintenanceAssetWeight uint32,
	initialLiabilityWeight uint32,
	maintenanceLiabilityWeight uint32,
	imfFactor uint32,
	liquidatorFee uint32,
	ifLiquidationFee uint32,
	activeStatus bool,
	assetTier AssetTier,
	scaleInitialAssetWeightStart uint64,
	withdrawGuardThreshold uint64,
	orderTickSize uint64,
	orderStepSize uint64,
	ifTotalFactor uint32,
	name [32]uint8,
	// Accounts:
	spotMarket ag_solanago.PublicKey,
	spotMarketMint ag_solanago.PublicKey,
	spotMarketVault ag_solanago.PublicKey,
	insuranceFundVault ag_solanago.PublicKey,
	driftSigner ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *InitializeSpotMarket {
	return NewInitializeSpotMarketInstructionBuilder().
		SetOptimalUtilization(optimalUtilization).
		SetOptimalBorrowRate(optimalBorrowRate).
		SetMaxBorrowRate(maxBorrowRate).
		SetOracleSource(oracleSource).
		SetInitialAssetWeight(initialAssetWeight).
		SetMaintenanceAssetWeight(maintenanceAssetWeight).
		SetInitialLiabilityWeight(initialLiabilityWeight).
		SetMaintenanceLiabilityWeight(maintenanceLiabilityWeight).
		SetImfFactor(imfFactor).
		SetLiquidatorFee(liquidatorFee).
		SetIfLiquidationFee(ifLiquidationFee).
		SetActiveStatus(activeStatus).
		SetAssetTier(assetTier).
		SetScaleInitialAssetWeightStart(scaleInitialAssetWeightStart).
		SetWithdrawGuardThreshold(withdrawGuardThreshold).
		SetOrderTickSize(orderTickSize).
		SetOrderStepSize(orderStepSize).
		SetIfTotalFactor(ifTotalFactor).
		SetName(name).
		SetSpotMarketAccount(spotMarket).
		SetSpotMarketMintAccount(spotMarketMint).
		SetSpotMarketVaultAccount(spotMarketVault).
		SetInsuranceFundVaultAccount(insuranceFundVault).
		SetDriftSignerAccount(driftSigner).
		SetStateAccount(state).
		SetOracleAccount(oracle).
		SetAdminAccount(admin).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}
