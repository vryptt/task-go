package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"system-monitor/controllers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	api := r.PathPrefix("/api/system").Subrouter()

	api.HandleFunc("/cpu", controllers.GetCPUInfo).Methods("GET")
	api.HandleFunc("/memory", controllers.GetMemoryInfo).Methods("GET")
	api.HandleFunc("/disk", controllers.GetDiskInfo).Methods("GET")
	api.HandleFunc("/network", controllers.GetNetworkInfo).Methods("GET")
	api.HandleFunc("/processes", controllers.GetProcessInfo).Methods("GET")
	api.HandleFunc("/info", controllers.GetSystemInfo).Methods("GET")
	api.HandleFunc("/gpu", controllers.GetGPUInfo).Methods("GET")
	api.HandleFunc("/containers", controllers.GetContainerInfo).Methods("GET")

	api.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		cpu := controllers.GetCPUInfoData()
		memory := controllers.GetMemoryInfoData()
		disk := controllers.GetDiskInfoData()
		network := controllers.GetNetworkInfoData()
		processes := controllers.GetProcessInfoData()
		info := controllers.GetSystemInfoData()
		gpu := controllers.GetGPUInfoData()
		containers := controllers.GetContainerInfoData()

		result := map[string]interface{}{
			"cpu":        cpu,
			"memory":     memory,
			"disk":       disk,
			"network":    network,
			"processes":  processes,
			"info":       info,
			"gpu":        gpu,
			"containers": containers,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}).Methods("GET")

	r.HandleFunc("/healthz", controllers.Healthz).Methods("GET")

	return r
}
