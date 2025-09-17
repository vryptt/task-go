package models

import "time"

type SystemInfo struct {
	Hostname      string     `json:"hostname"`
	UptimeSeconds uint64     `json:"uptimeSeconds"`
	OS            OSInfo     `json:"os"`
	Users         []UserInfo `json:"users"`
}

type OSInfo struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Kernel       string `json:"kernel"`
	Architecture string `json:"architecture"`
}

type UserInfo struct {
	Username  string    `json:"username"`
	TTY       string    `json:"tty"`
	LoginTime time.Time `json:"loginTime"`
}
