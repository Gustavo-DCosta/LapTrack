package router

import (
	"github.com/Gustavo-DCosta/LapTrack/services"
	"net/http"
)

func receiveTelemetry(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := services.TelemetryDecoder(r)
	if err != nil {
		http.Error(w, "Invalid JSON, bad request", http.StatusBadRequest)
		return
	}

	_ = data

	w.WriteHeader(http.StatusOK)
}
func Routing() {
	http.HandleFunc("/api/telemetry", receiveTelemetry)
}