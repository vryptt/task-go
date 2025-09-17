package controllers

import (
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"system-monitor/models"
)

func GetSystemInfoData() models.SystemInfo {
	info, err := host.Info()
	if err != nil {
		return models.SystemInfo{}
	}

	rawUsers, _ := host.Users()
	users := []models.UserInfo{}
	for _, u := range rawUsers {
		users = append(users, models.UserInfo{
			Username:  u.User,
			TTY:       u.Terminal,
			LoginTime: time.Unix(int64(u.Started), 0),
		})
	}

	return models.SystemInfo{
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
}

func GetSystemInfo(w http.ResponseWriter, r *http.Request) {
	data := GetSystemInfoData()
	if data.Hostname == "" {
		RespondJSON(w, JSONResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to fetch system info",
		})
		return
	}

	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: data,
	})
}