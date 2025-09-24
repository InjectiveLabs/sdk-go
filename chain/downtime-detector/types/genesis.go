package types

import "time"

func DefaultGenesis() *GenesisState {
	genDowntimes := []GenesisDowntimeEntry{}
	for _, downtime := range DowntimeToDuration.Keys() {
		genDowntimes = append(genDowntimes, GenesisDowntimeEntry{
			Duration:     downtime,
			LastDowntime: DefaultLastDowntime,
		})
	}
	return &GenesisState{
		Downtimes:     genDowntimes,
		LastBlockTime: time.Unix(0, 0),
	}
}

func (*GenesisState) Validate() error {
	return nil
}

func NewGenesisDowntimeEntry(dur Downtime, tm time.Time) GenesisDowntimeEntry {
	return GenesisDowntimeEntry{Duration: dur, LastDowntime: tm}
}
