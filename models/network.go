package models

type NetworkInfo struct {
	Interfaces  []NetworkInterface  `json:"interfaces"`
	Connections []NetworkConnection `json:"connections"`
	IOStats     NetworkIOMap        `json:"ioStats,omitempty"`
}

type NetworkInterface struct {
	Name          string `json:"name"`
	IPv4          string `json:"ipv4"`
	RxBytes       uint64 `json:"rxBytes"`
	TxBytes       uint64 `json:"txBytes"`
	RxPackets     uint64 `json:"rxPackets"`
	TxPackets     uint64 `json:"txPackets"`
	Errors        uint64 `json:"errors"`
	Dropped       uint64 `json:"dropped"`
}

type NetworkConnection struct {
	Protocol      string `json:"protocol"`
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
	State         string `json:"state"`
	PID           int32  `json:"pid"`
	Process       string `json:"process"`
}

type NetworkIO struct {
	RxBytesPerSec float64 `json:"rxBytesPerSec"`
	TxBytesPerSec float64 `json:"txBytesPerSec"`
}

type NetworkIOMap map[string]NetworkIO
