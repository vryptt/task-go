package controllers

import (
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"system-monitor/models"
)

func GetProcessInfo(w http.ResponseWriter, r *http.Request) {
	procs, err := process.Processes()
	if err != nil {
		RespondJSON(w, JSONResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	var running, sleeping, stopped, zombie int
	var top []models.ProcessDetail

	for _, p := range procs {
		status, _ := p.Status()
		if len(status) > 0 {
			switch status[0] {
			case "R":
				running++
			case "S":
				sleeping++
			case "T":
				stopped++
			case "Z":
				zombie++
			}
		}
		if len(top) < 15 {
			name, _ := p.Name()
			user, _ := p.Username()
			cpuPerc, _ := p.CPUPercent()
			memPerc, _ := p.MemoryPercent()
			threads, _ := p.NumThreads()
			start, _ := p.CreateTime()
			top = append(top, models.ProcessDetail{
				PID:        p.Pid,
				User:       user,
				Name:       name,
				CPUPercent: cpuPerc,
				MemPercent: memPerc,
				Threads:    threads,
				StartTime:  time.Unix(0, start*int64(time.Millisecond)),
			})
		}
	}

	RespondJSON(w, JSONResponse{
		Status: http.StatusOK,
		Payload: models.ProcessInfo{
			Total:    len(procs),
			Running:  running,
			Sleeping: sleeping,
			Stopped:  stopped,
			Zombie:   zombie,
			Top:      top,
		},
	})
}