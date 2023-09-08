package types

import (
	fnfttypes "github.com/FluxNFTLabs/sdk-go/chain/modules/fnft/types"
)

// ValidateGenesis checks that the given genesis state has no integrity issues
func ValidateGenesis(data GenesisState) error {
	for _, product := range data.Products {
		if len(product.ClassId) == 0 {
			return fnfttypes.ErrEmptyClassID
		}
		if len(product.Id) == 0 {
			return fnfttypes.ErrEmptyNFTID
		}
		if len(product.ProductId) == 0 {
			return ErrEmptyProductID
		}
		if len(product.Offerings) == 0 {
			return ErrEmptyOfferings
		}
	}
	for _, c := range data.Commissions {
		if len(c.ClassId) == 0 {
			return fnfttypes.ErrEmptyClassID
		}
		if c.CommissionMul == 0 || c.CommissionDiv == 0 {
			return ErrInvalidCommissionPart
		}
	}
	return nil
}

// DefaultGenesisState - Returns a default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Products: []*Product{},
		Commissions: []*ClassCommission{
			{
				ClassId:       "series",
				CommissionMul: 15,
				CommissionDiv: 100,
			},
			{
				ClassId:       "livestream",
				CommissionMul: 50,
				CommissionDiv: 100,
			},
			{
				ClassId:       "music",
				CommissionMul: 15,
				CommissionDiv: 100,
			},
			{
				ClassId:       "consuming",
				CommissionMul: 15,
				CommissionDiv: 100,
			},
		},
	}
}
