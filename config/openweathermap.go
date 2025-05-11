package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log/slog"
)

type OpenWeatherMapConfiguration struct {
	OpenWeatherMapEndPointURL string `env:"OPEN_WEATHER_MAP_ENDPOINT_URL, default=https://api.openweathermap.org"`
	AppID                     string `env:"OPEN_WEATHER_MAP_APP_ID, required"`
	Lat                       string `env:"OPEN_WEATHER_MAP_LAT, required"`
	Lon                       string `env:"OPEN_WEATHER_MAP_LON, required"`
}

func NewOpenWeatherMapConfiguration(ctx context.Context) *OpenWeatherMapConfiguration {
	var c OpenWeatherMapConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration on NewOpenWeatherMapConfiguration", "err", err)
		panic(err)
	}

	return &c
}
