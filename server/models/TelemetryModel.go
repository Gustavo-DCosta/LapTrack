package models

type TelemetryModel struct {
	Track string `json:"track"`
	Car string `json:"car"`
	Weather string `json:"weather"`
	LapTimes []string `json:"lap_times"`
}