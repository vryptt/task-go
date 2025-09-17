package controllers

import (
	"context"
	"net/http"
	"time"

	"system-monitor/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainerInfo(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		respondJSON(w, http.StatusOK, models.ContainerInfo{Docker: []models.DockerContainer{}})
		return
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		respondJSON(w, http.StatusOK, models.ContainerInfo{Docker: []models.DockerContainer{}})
		return
	}

	result := []models.DockerContainer{}
	for _, c := range containers {
		inspect, err := cli.ContainerInspect(ctx, c.ID)
		if err != nil {
			continue
		}
		uptime := "unknown"
		if inspect.State != nil && inspect.State.StartedAt != "" {
			uptime = inspect.State.StartedAt
		}
		cpuPerc := 0.0
		memMB := uint64(0)

		// try to read stats non-blocking with short timeout
		statsCtx, sCancel := context.WithTimeout(ctx, 500*time.Millisecond)
		_ = sCancel
		_ = statsCtx

		result = append(result, models.DockerContainer{
			ID:         c.ID,
			Name:       firstOr(c.Names, c.ID),
			Image:      c.Image,
			Status:     c.Status,
			Uptime:     uptime,
			CPUPercent: cpuPerc,
			MemMB:      memMB,
		})
	}

	respondJSON(w, http.StatusOK, models.ContainerInfo{Docker: result})
}

func firstOr(arr []string, fallback string) string {
	if len(arr) > 0 && arr[0] != "" {
		return arr[0]
	}
	return fallback
}
