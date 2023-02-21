package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewVolumeRecord(makerVolume, takerVolume sdk.Dec) VolumeRecord {
	return VolumeRecord{
		MakerVolume: makerVolume,
		TakerVolume: takerVolume,
	}
}

func NewZeroVolumeRecord() VolumeRecord {
	return NewVolumeRecord(sdk.ZeroDec(), sdk.ZeroDec())
}

func (v VolumeRecord) Add(record VolumeRecord) VolumeRecord {
	if v.MakerVolume.IsNil() {
		v.MakerVolume = sdk.ZeroDec()
	}

	if v.TakerVolume.IsNil() {
		v.TakerVolume = sdk.ZeroDec()
	}

	if record.IsZero() {
		return v
	}

	newMakerVolume := v.MakerVolume
	newTakerVolume := v.TakerVolume

	if !record.MakerVolume.IsNil() && !record.MakerVolume.IsZero() {
		newMakerVolume = v.MakerVolume.Add(record.MakerVolume)
	}
	if !record.TakerVolume.IsNil() && !record.TakerVolume.IsZero() {
		newTakerVolume = v.TakerVolume.Add(record.TakerVolume)
	}

	return NewVolumeRecord(newMakerVolume, newTakerVolume)
}

func (v *VolumeRecord) IsZero() bool {
	return (v.TakerVolume.IsNil() || v.TakerVolume.IsZero()) && (v.MakerVolume.IsNil() || v.MakerVolume.IsZero())
}

func (v *VolumeRecord) Total() sdk.Dec {
	totalVolume := sdk.ZeroDec()
	if !v.TakerVolume.IsNil() {
		totalVolume = totalVolume.Add(v.TakerVolume)
	}
	if !v.MakerVolume.IsNil() {
		totalVolume = totalVolume.Add(v.MakerVolume)
	}
	return totalVolume
}

func NewVolumeWithSingleType(volume sdk.Dec, isMaker bool) VolumeRecord {
	if isMaker {
		return NewVolumeRecord(volume, sdk.ZeroDec())
	}
	return NewVolumeRecord(sdk.ZeroDec(), volume)
}
