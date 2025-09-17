package controllers

import (
	"github.com/shirou/gopsutil/v3/host"
)

func GetCPUTemperaturesData() []float64 {
	temps, err := host.SensorsTemperatures()
	if err != nil {
		return []float64{}
	}
	out := []float64{}
	for _, t := range temps {
		if t.SensorKey == "Package id 0" || 
		   t.SensorKey == "Tdie" || 
		   t.SensorKey == "CPU Temperature" || 
		   t.SensorKey == "cpu_thermal" || 
		   t.SensorKey == "coretemp" {
			out = append(out, t.Temperature)
		} else if t.Temperature > 0 {
			out = append(out, t.Temperature)
		}
	}
	return out
}