package perf

import (
	"api/src/core/base"
	"api/src/db"
	"fmt"
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type service struct {
	repositoryManager *db.RepositoryManager
	logger            *base.Logger
}

func newService(repositoryManager *db.RepositoryManager, logger *base.Logger) *service {
	sLogger := logger.With("service", "perf")
	return &service{
		repositoryManager: repositoryManager,
		logger:            sLogger,
	}
}

func (s *service) getPerf() (*perfDetails, error) {
	// --- 1. Charge CPU (%) ---
	percent, err := cpu.Percent(0, false) // false -> moyenne globale
	if err != nil {
		return nil, fmt.Errorf("err")
	}
	cpuLoad := percent[0]

	// --- 2. Température CPU (°C) ---
	var tmp float64 = -1
	temps, err := host.SensorsTemperatures()
	if err != nil {
		log.Printf("Erreur host.SensorsTemperatures: %v", err)
	} else {
		for _, t := range temps {
			if t.SensorKey == "Package id 0" || t.SensorKey == "CPU" {
				tmp = t.Temperature
			}
		}
	}

	// --- 3. Mémoire utilisée / totale ---
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("erreur")
	}

	// --- 4. Estimation puissance CPU (W) ---
	const TDP_CPU = 45.0
	cpuPower := TDP_CPU * (cpuLoad / 100)
	return newPerfDetails(cpuLoad, tmp, cpuPower, float64(vmStat.Used)/1024/1024/1024, float64(vmStat.Total)/1024/1024/1024), nil
}
