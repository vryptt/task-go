package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	psnet "github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"system-monitor/models"
)

var lastNetIO = struct {
	ts   time.Time
	data map[string]psnet.IOCountersStat
}{data: map[string]psnet.IOCountersStat{}}

func GetNetworkInfo(w http.ResponseWriter, r *http.Request) {
	ifaces, err := psnet.Interfaces()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	var interfaces []models.NetworkInterface
	for _, iface := range ifaces {
		ipv4 := ""
		for _, addr := range iface.Addrs {
			if addr.Addr != "" && !strings.Contains(addr.Addr, ":") {
				ipv4 = strings.Split(addr.Addr, "/")[0]
				break
			}
		}
		interfaces = append(interfaces, models.NetworkInterface{
			Name:      iface.Name,
			IPv4:      ipv4,
			RxBytes:   iface.BytesRecv,
			TxBytes:   iface.BytesSent,
			RxPackets: iface.PacketsRecv,
			TxPackets: iface.PacketsSent,
			Errors:    iface.Errin + iface.Errout,
			Dropped:   iface.Dropin + iface.Dropout,
		})
	}

	conns, _ := psnet.Connections("all")
	var connections []models.NetworkConnection
	for i, c := range conns {
		if i >= 20 {
			break
		}
		procName := "unknown"
		if c.Pid != 0 {
			if proc, err := process.NewProcess(c.Pid); err == nil {
				if name, err := proc.Name(); err == nil {
					procName = name
				}
			}
		}
		proto := fmt.Sprintf("%d", c.Type)
		if c.Type == 1 {
			proto = "tcp"
		} else if c.Type == 2 {
			proto = "udp"
		}
		connections = append(connections, models.NetworkConnection{
			Protocol:      proto,
			LocalAddress:  fmt.Sprintf("%s:%d", c.Laddr.IP, c.Laddr.Port),
			RemoteAddress: fmt.Sprintf("%s:%d", c.Raddr.IP, c.Raddr.Port),
			State:         c.Status,
			PID:           c.Pid,
			Process:       procName,
		})
	}

	// IO per-second (delta sampling)
	ios, _ := psnet.IOCounters(true)
	now := time.Now()
	elapsed := now.Sub(lastNetIO.ts).Seconds()
	if elapsed <= 0 {
		elapsed = 1
	}
	ioMap := models.NetworkIOMap{}
	for _, stat := range ios {
		prev, ok := lastNetIO.data[stat.Name]
		var entry models.NetworkIO
		if ok {
			entry.RxBytesPerSec = float64(stat.BytesRecv-prev.BytesRecv) / elapsed
			entry.TxBytesPerSec = float64(stat.BytesSent-prev.BytesSent) / elapsed
		}
		ioMap[stat.Name] = entry
	}

	lastNetIO.ts = now
	// rebuild map
	m := map[string]psnet.IOCountersStat{}
	for _, s := range ios {
		m[s.Name] = s
	}
	lastNetIO.data = m

	respondJSON(w, http.StatusOK, models.NetworkInfo{
		Interfaces:  interfaces,
		Connections: connections,
		IOStats:     ioMap,
	})
}
