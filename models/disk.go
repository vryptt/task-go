package models

type DiskInfo struct {
	Partitions []Partition `json:"partitions"`
	IOStats    DiskIOMap   `json:"ioStats,omitempty"`
}

type Partition struct {
	Mount        string  `json:"mount"`
	Filesystem   string  `json:"filesystem"`
	TotalGB      float64 `json:"totalGB"`
	UsedGB       float64 `json:"usedGB"`
	FreeGB       float64 `json:"freeGB"`
	UsagePercent float64 `json:"usagePercent"`
}

type DiskIO struct {
	ReadBytesPerSec  float64 `json:"readBytesPerSec"`
	WriteBytesPerSec float64 `json:"writeBytesPerSec"`
	ReadOpsPerSec    float64 `json:"readOpsPerSec"`
	WriteOpsPerSec   float64 `json:"writeOpsPerSec"`
}

type DiskIOMap map[string]DiskIO
