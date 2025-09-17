package models

type CPUInfo struct {
	Usage        CPUUsage `json:"usage"`
	LoadAverage  LoadAvg  `json:"loadAverage"`
	FrequencyMHz float64  `json:"frequencyMHz"`
	TemperatureC float64  `json:"temperatureC"`
	Cores        int      `json:"cores"`
}

type CPUUsage struct {
	Total   float64   `json:"total"`
	PerCore []float64 `json:"perCore"`
}

type LoadAvg struct {
	OneMin     float64 `json:"1m"`
	FiveMin    float64 `json:"5m"`
	FifteenMin float64 `json:"15m"`
}
