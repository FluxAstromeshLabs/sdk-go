package types

const (
	// ModuleName defines the name of the svm module
	ModuleName = "interpool"

	// StoreKey is the default store key for svm
	StoreKey = ModuleName

	// RouterKey is the message route for svm
	RouterKey = ModuleName

	FeeRatePrecision    = 1_000_000
	PercentagePrecision = 10_000

	DefaultPoolSvmLamport = 1_000_000_000
)
