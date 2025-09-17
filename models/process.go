package models

import "time"

type ProcessInfo struct {
	Total    int             `json:"total"`
	Running  int             `json:"running"`
	Sleeping int             `json:"sleeping"`
	Stopped  int             `json:"stopped"`
	Zombie   int             `json:"zombie"`
	Top      []ProcessDetail `json:"top"`
}

type ProcessDetail struct {
	PID        int32     `json:"pid"`
	User       string    `json:"user"`
	Name       string    `json:"name"`
	CPUPercent float64   `json:"cpuPercent"`
	MemPercent float32   `json:"memPercent"`
	Threads    int32     `json:"threads"`
	StartTime  time.Time `json:"startTime"`
}
