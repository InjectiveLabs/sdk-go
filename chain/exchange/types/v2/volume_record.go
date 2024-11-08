package v2

import (
	"cosmossdk.io/math"
)

func NewVolumeRecord(makerVolume, takerVolume math.LegacyDec) VolumeRecord {
	return VolumeRecord{
		MakerVolume: makerVolume,
		TakerVolume: takerVolume,
	}
}

func NewZeroVolumeRecord() VolumeRecord {
	return NewVolumeRecord(math.LegacyZeroDec(), math.LegacyZeroDec())
}

func (v VolumeRecord) Add(record VolumeRecord) VolumeRecord {
	if v.MakerVolume.IsNil() {
		v.MakerVolume = math.LegacyZeroDec()
	}

	if v.TakerVolume.IsNil() {
		v.TakerVolume = math.LegacyZeroDec()
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

func (v *VolumeRecord) Total() math.LegacyDec {
	totalVolume := math.LegacyZeroDec()
	if !v.TakerVolume.IsNil() {
		totalVolume = totalVolume.Add(v.TakerVolume)
	}
	if !v.MakerVolume.IsNil() {
		totalVolume = totalVolume.Add(v.MakerVolume)
	}
	return totalVolume
}

func NewVolumeWithSingleType(volume math.LegacyDec, isMaker bool) VolumeRecord {
	if isMaker {
		return NewVolumeRecord(volume, math.LegacyZeroDec())
	}
	return NewVolumeRecord(math.LegacyZeroDec(), volume)
}
