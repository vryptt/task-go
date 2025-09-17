package controllers

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"net/http"
	"strconv"

	"system-monitor/models"
)

func GetGPUInfo(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("nvidia-smi",
		"--query-gpu=name,driver_version,memory.total,memory.used,utilization.gpu,temperature.gpu",
		"--format=csv,noheader,nounits").Output()
	if err != nil {
		RespondJSON(w, JSONResponse{
			Status:  http.StatusOK,
			Payload: models.GPUInfo{GPUs: []models.GPU{}},
		})
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	gpus := []models.GPU{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) < 6 {
			continue
		}
		memTotal, _ := strconv.ParseUint(parts[2], 10, 64)
		memUsed, _ := strconv.ParseUint(parts[3], 10, 64)
		util, _ := strconv.ParseFloat(parts[4], 64)
		temp, _ := strconv.ParseFloat(parts[5], 64)

		gpus = append(gpus, models.GPU{
			Vendor:          "NVIDIA",
			Name:            parts[0],
			Driver:          parts[1],
			MemoryTotalMB:   memTotal,
			MemoryUsedMB:    memUsed,
			GPUUsagePercent: util,
			TemperatureC:    temp,
		})
	}

	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: models.GPUInfo{GPUs: gpus},
	})
}