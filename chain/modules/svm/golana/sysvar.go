package golana

/*
#cgo LDFLAGS: -L./../lib -lgolana
#include "./../lib/golana.h"
*/
import "C"

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/FluxNFTLabs/sdk-go/chain/modules/svm/types"
	bin "github.com/gagliardetto/binary"
)

type Rent struct {
	LamportsPerByteYear uint64
	ExemptionPeriod     float64
	BurnPercent         byte
}

var DefaultRent = Rent{
	LamportsPerByteYear: types.DefaultLamportsPerByteYear,
	ExemptionPeriod:     types.DefaultExemptionThreshold,
	BurnPercent:         types.DefaultBurnPercent,
}

type EpochSchedule struct {
	SlotsPerEpoch            uint64
	LeaderScheduleSlotOffset uint64
	WarmUp                   bool
	FirstNormalEpoch         uint64
	FirstNormalSlot          uint64
}

var DefaultEpochSchedule = EpochSchedule{
	SlotsPerEpoch:            types.DefaultSlotsPerEpoch,
	LeaderScheduleSlotOffset: types.DefaultLeaderScheduleSlotOffset,
	WarmUp:                   true,
}

type Clock struct {
	slot                uint64
	epochStartTimestamp int64
	epoch               uint64
	leaderScheduleEpoch uint64
	unixTimestamp       int64
}

func NewClock(ctx sdk.Context) Clock {
	blockTime := ctx.BlockTime().Unix()
	slot := uint64(ctx.BlockHeight())
	epoch := slot / DEFAULT_SLOTS_PER_EPOCH
	if slot%DEFAULT_SLOTS_PER_EPOCH == 0 {
		epochStartTs = blockTime
	}
	return Clock{
		slot:                slot,
		epochStartTimestamp: epochStartTs,
		epoch:               epoch,
		leaderScheduleEpoch: epoch,
		unixTimestamp:       blockTime,
	}
}

func MarshalBinary(x interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	enc := bin.NewBinEncoder(buffer)
	if err := enc.Encode(x); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func UnmarshalBinary(buffer []byte, obj interface{}) error {
	enc := bin.NewBinDecoder(buffer)
	err := enc.Decode(obj)
	if err != nil {
		return err
	}

	return nil
}

type SysvarCache struct {
	sysvarCache *C.golana_sysvar_cache
}

func NewSysvarCache() SysvarCache {
	return SysvarCache{
		sysvarCache: C.golana_sysvar_cache_new(),
	}
}

func (sc SysvarCache) SetClock(v SysvarClock) {
	C.golana_sysvar_cache_set_clock(sc.sysvarCache, v.clock)
}

func (sc SysvarCache) SetRent(v SysvarRent) {
	C.golana_sysvar_cache_set_rent(sc.sysvarCache, v.rent)
}

func FreeSysvarCache(sc SysvarCache) {
	C.golana_sysvar_cache_free(sc.sysvarCache)
}

type SysvarClock struct {
	clock *C.golana_sysvar_clock
}

func NewSysvarClock(clock Clock) SysvarClock {
	c := C.golana_sysvar_clock_new(
		C.uint64_t(clock.slot),
		C.int64_t(clock.epochStartTimestamp),
		C.uint64_t(clock.epoch),
		C.uint64_t(clock.leaderScheduleEpoch),
		C.int64_t(clock.unixTimestamp),
	)
	return SysvarClock{clock: c}
}

func FreeSysvarClock(v SysvarClock) {
	C.golana_sysvar_clock_free(v.clock)
}

type SysvarRent struct {
	rent *C.golana_sysvar_rent
}

func NewSysvarRent(
	lamportsPerByteYear uint64,
	exemptionThreshold float64,
	burnPercent byte,
) SysvarRent {
	return SysvarRent{
		rent: C.golana_sysvar_rent_new(
			C.ulonglong(lamportsPerByteYear),
			C.double(exemptionThreshold),
			C.uchar(burnPercent),
		),
	}
}

func FreeSysvarRent(v SysvarRent) {
	C.golana_sysvar_rent_free(v.rent)
}
