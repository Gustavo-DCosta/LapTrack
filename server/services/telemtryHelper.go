package services

import (
	"encoding/json"
	"github.com/Gustavo-DCosta/LapTrack/models"
	"net/http"
)

func TelemetryDecoder(r *http.Request) (models.TelemetryModel, error) {
	var t models.TelemetryModel

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&t); if err != nil {
		return t, err
	}

	return t, nil
}


func ProcessTelemetry(model models.TelemetryModel) (bool){
	var isValid bool
	var t models.TelemetryModel

	if t.Car == "" || t.Track == "" || t.Weather == "" {
		isValid = false
	}
	if t.LapTimes == nil {
		isValid = false
	}

	
	return isValid
}