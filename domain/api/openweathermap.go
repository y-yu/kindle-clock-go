package api

import "context"

type OpenWeatherMapInfo struct {
	Weather []struct {
		Icon string `json:"icon" validate:"required"`
	} `json:"weather" validate:"required"`
}

type OpenWeatherMapAPIClient interface {
	GetLatest(ctx context.Context) (OpenWeatherMapInfo, error)
}
