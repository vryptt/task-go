package models

type GPUInfo struct {
	GPUs []GPU `json:"gpus"`
}

type GPU struct {
	Vendor          string  `json:"vendor"`
	Name            string  `json:"name"`
	Driver          string  `json:"driver"`
	MemoryTotalMB   uint64  `json:"memoryTotalMB"`
	MemoryUsedMB    uint64  `json:"memoryUsedMB"`
	GPUUsagePercent float64 `json:"gpuUsagePercent"`
	TemperatureC    float64 `json:"temperatureC"`
}
