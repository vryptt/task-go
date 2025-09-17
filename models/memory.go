package models

type MemoryInfo struct {
	RAM  RAMInfo  `json:"ram"`
	Swap SwapInfo `json:"swap"`
}

type RAMInfo struct {
	TotalMB      uint64  `json:"totalMB"`
	UsedMB       uint64  `json:"usedMB"`
	FreeMB       uint64  `json:"freeMB"`
	CachedMB     uint64  `json:"cachedMB"`
	BuffersMB    uint64  `json:"buffersMB"`
	UsagePercent float64 `json:"usagePercent"`
}

type SwapInfo struct {
	TotalMB      uint64  `json:"totalMB"`
	UsedMB       uint64  `json:"usedMB"`
	FreeMB       uint64  `json:"freeMB"`
	UsagePercent float64 `json:"usagePercent"`
}
