package controllers

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status  int         `json:"-"`
	Payload interface{} `json:"payload,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondJSON(w http.ResponseWriter, resp JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)

	if resp.Error != "" {
		_ = json.NewEncoder(w).Encode(map[string]string{"error": resp.Error})
		return
	}

	if resp.Payload != nil {
		_ = json.NewEncoder(w).Encode(resp.Payload)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]any{})
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, JSONResponse{
		Status:  http.StatusOK,
		Payload: map[string]string{"status": "ok"},
	})
}