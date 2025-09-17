package models

type ContainerInfo struct {
	Docker []DockerContainer `json:"docker"`
}

type DockerContainer struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	Status     string  `json:"status"`
	Uptime     string  `json:"uptime"`
	CPUPercent float64 `json:"cpuPercent"`
	MemMB      uint64  `json:"memMB"`
}
