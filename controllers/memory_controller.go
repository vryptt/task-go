package controllers

import (
	"net/http"

	"github.com/shirou/gopsutil/v3/mem"
	"system-monitor/models"
	"system-monitor/utils"
)

func GetMemoryInfo(w http.ResponseWriter, r *http.Request) {
	data := GetMemoryInfoData()
	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: data,
	})
}

func GetMemoryInfoData() models.MemoryInfo {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return models.MemoryInfo{}
	}
	swap, _ := mem.SwapMemory()

	return models.MemoryInfo{
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
}