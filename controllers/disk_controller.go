package controllers

import (
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
	"system-monitor/models"
	"system-monitor/utils"
)

// sampling state
var lastDiskIO = struct {
	ts   time.Time
	data map[string]disk.IOCountersStat
}{data: map[string]disk.IOCountersStat{}}

func GetDiskInfo(w http.ResponseWriter, r *http.Request) {
	parts, err := disk.Partitions(false)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	var partitions []models.Partition
	for _, p := range parts {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		partitions = append(partitions, models.Partition{
			Mount:        p.Mountpoint,
			Filesystem:   p.Fstype,
			TotalGB:      utils.BytesToGB(usage.Total),
			UsedGB:       utils.BytesToGB(usage.Used),
			FreeGB:       utils.BytesToGB(usage.Free),
			UsagePercent: usage.UsedPercent,
		})
	}

	ioMap := models.DiskIOMap{}
	ios, _ := disk.IOCounters()
	now := time.Now()
	elapsed := now.Sub(lastDiskIO.ts).Seconds()
	if elapsed <= 0 {
		elapsed = 1
	}
	for name, stat := range ios {
		prev, ok := lastDiskIO.data[name]
		var d models.DiskIO
		if ok {
			d.ReadBytesPerSec = float64(stat.ReadBytes-prev.ReadBytes) / elapsed
			d.WriteBytesPerSec = float64(stat.WriteBytes-prev.WriteBytes) / elapsed
			d.ReadOpsPerSec = float64(stat.ReadCount-prev.ReadCount) / elapsed
			d.WriteOpsPerSec = float64(stat.WriteCount-prev.WriteCount) / elapsed
		}
		ioMap[name] = d
	}

	// update last
	lastDiskIO.ts = now
	lastDiskIO.data = ios

	respondJSON(w, http.StatusOK, models.DiskInfo{Partitions: partitions, IOStats: ioMap})
}
