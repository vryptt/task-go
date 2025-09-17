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
	data := GetContainerInfoData()
	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: data,
	})
}

func GetContainerInfoData() models.ContainerInfo {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return models.ContainerInfo{Docker: []models.DockerContainer{}}
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return models.ContainerInfo{Docker: []models.DockerContainer{}}
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

		result = append(result, models.DockerContainer{
			ID:         c.ID,
			Name:       firstOr(c.Names, c.ID),
			Image:      c.Image,
			Status:     c.Status,
			Uptime:     uptime,
			CPUPercent: 0.0,
			MemMB:      0,
		})
	}

	return models.ContainerInfo{Docker: result}
}

func firstOr(arr []string, fallback string) string {
	if len(arr) > 0 && arr[0] != "" {
		return arr[0]
	}
	return fallback
}