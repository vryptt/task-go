package routes

import (
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

	// simple health
	r.HandleFunc("/healthz", controllers.Healthz).Methods("GET")

	return r
}
