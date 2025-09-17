package controllers

import (
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"system-monitor/models"
)

func GetSystemInfo(w http.ResponseWriter, r *http.Request) {
	info, err := host.Info()
	if err != nil {
		RespondJSON(w, JSONResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	usersRaw, _ := host.Users()
	users := []models.UserInfo{}
	for _, u := range usersRaw {
		users = append(users, models.UserInfo{
			Username:  u.User,
			TTY:       u.Terminal,
			LoginTime: time.Unix(int64(u.Started), 0),
		})
	}

	data := models.SystemInfo{
		Hostname:      info.Hostname,
		UptimeSeconds: info.Uptime,
		OS: models.OSInfo{
			Name:         info.Platform,
			Version:      info.PlatformVersion,
			Kernel:       info.KernelVersion,
			Architecture: info.KernelArch,
		},
		Users: users,
	}

	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: data,
	})
}