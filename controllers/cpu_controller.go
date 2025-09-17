package controllers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"system-monitor/models"
)

func GetCPUInfo(w http.ResponseWriter, r *http.Request) {
	perCore, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		RespondJSON(w, JSONResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}
	total, err := cpu.Percent(0, false)
	if err != nil {
		RespondJSON(w, JSONResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}
	loadAvg, _ := load.Avg()

	info, _ := cpu.Info()
	var mhz float64
	if len(info) > 0 {
		mhz = info[0].Mhz
	}

	temp := 0.0
	if temps, err := cpuTemperatures(); err == nil && len(temps) > 0 {
		temp = temps[0]
	}

	data := models.CPUInfo{
		Usage: models.CPUUsage{
			Total:   0,
			PerCore: perCore,
		},
		LoadAverage: models.LoadAvg{
			OneMin:     loadAvg.Load1,
			FiveMin:    loadAvg.Load5,
			FifteenMin: loadAvg.Load15,
		},
		FrequencyMHz: mhz,
		TemperatureC: temp,
		Cores:        runtime.NumCPU(),
	}

	if len(total) > 0 {
		data.Usage.Total = total[0]
	} else {
		var sum float64
		for _, v := range perCore {
			sum += v
		}
		if len(perCore) > 0 {
			data.Usage.Total = sum / float64(len(perCore))
		}
	}

	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: data,
	})
}