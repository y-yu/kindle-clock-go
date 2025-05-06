package repository

import (
	"context"
	"github.com/y-yu/kindle-clock-go/domain/model"
)

type OpenWeatherMapRepository interface {
	GetCurrentWeather(ctx context.Context) (model.Weather, error)
}
