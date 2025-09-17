package controllers

import (
	"net/http"

	"github.com/shirou/gopsutil/v3/mem"
	"system-monitor/models"
	"system-monitor/utils"
)

func GetMemoryInfo(w http.ResponseWriter, r *http.Request) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		RespondJSON(w, JSONResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}
	swap, _ := mem.SwapMemory()

	data := models.MemoryInfo{
		RAM: models.RAMInfo{
			TotalMB:      utils.BytesToMB(vm.Total),
			UsedMB:       utils.BytesToMB(vm.Used),
			FreeMB:       utils.BytesToMB(vm.Free),
			CachedMB:     utils.BytesToMB(vm.Cached),
			BuffersMB:    utils.BytesToMB(vm.Buffers),
			UsagePercent: vm.UsedPercent,
		},
		Swap: models.SwapInfo{
			TotalMB:      utils.BytesToMB(swap.Total),
			UsedMB:       utils.BytesToMB(swap.Used),
			FreeMB:       utils.BytesToMB(swap.Free),
			UsagePercent: swap.UsedPercent,
		},
	}

	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: data,
	})
}