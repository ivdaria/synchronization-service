package model

import "time"

type AddClientRequestBody struct {
	ClientName  string    `json:"client_name"`
	Version     int       `json:"version"`
	Image       string    `json:"image"`
	CPU         string    `json:"cpu"`
	Memory      string    `json:"memory"`
	Priority    float64   `json:"priority"`
	NeedRestart bool      `json:"need_status"`
	SpawnedAt   time.Time `json:"spawned_at"`
}

type AddClientResponseBody struct {
	ID int64 `json:"id"`
}

type UpdateClientRequestBody struct {
	ClientName  string    `json:"client_name"`
	Version     int       `json:"version"`
	Image       string    `json:"image"`
	CPU         string    `json:"cpu"`
	Memory      string    `json:"memory"`
	Priority    float64   `json:"priority"`
	NeedRestart bool      `json:"need_status"`
	SpawnedAt   time.Time `json:"spawned_at"`
}

type UpdateStatusRequestBody struct {
	VWAP bool `json:"vwap"`
	TWAP bool `json:"twap"`
	HFT  bool `json:"hft"`
}
