package controllers

import (
	"github.com/shirou/gopsutil/v3/host"
)

// cpuTemperatures tries to read sensor temps; returns slice of floats (Celsius)
func cpuTemperatures() ([]float64, error) {
	temps, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	out := []float64{}
	for _, t := range temps {
		// filter common labels
		if t.SensorKey == "Package id 0" || t.SensorKey == "Tdie" || t.SensorKey == "CPU Temperature" || t.SensorKey == "cpu_thermal" || t.SensorKey == "coretemp" {
			out = append(out, t.Temperature)
		} else {
			if t.Temperature > 0 {
				out = append(out, t.Temperature)
			}
		}
	}
	return out, nil
}
