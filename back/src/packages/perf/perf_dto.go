package perf

type perfDetails struct {
	CPUCharge float64 `json:"cpu_charge"`
	CPUTemp   float64 `json:"cpu_temp"`
	CPUPower  float64 `json:"cpu_power"`
	UsedRam   float64 `json:"used_ram"`
	DispoRam  float64 `json:"dispo_ram"`
}

func newPerfDetails(cpuCharge, cpuTemp, cpuPower, userRam, dispoRam float64) *perfDetails {
	return &perfDetails{
		CPUCharge: cpuCharge,
		CPUTemp:   cpuTemp,
		CPUPower:  cpuPower,
		UsedRam:   userRam,
		DispoRam:  dispoRam,
	}
}
