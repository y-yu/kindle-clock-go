//go:build wireinject
// +build wireinject

package inject

import (
	"context"
	"github.com/google/wire"
	"github.com/y-yu/kindle-clock-go/domain"
	domainRepository "github.com/y-yu/kindle-clock-go/domain/repository"
	"github.com/y-yu/kindle-clock-go/infra/api"
	"github.com/y-yu/kindle-clock-go/infra/cache"
	"github.com/y-yu/kindle-clock-go/repository"
)

var binding = wire.NewSet(
	domain.NewSystemClock,
	api.NewAwairAPIClient,
	api.NewNatureRemoAPIClient,
	api.NewSwitchBotAPIClient,
	api.NewOpenWeatherMapAPIClient,
	cache.NewAwairCacheClient,
	cache.NewSwitchBotCacheClient,
	repository.NewAwairRepository,
	repository.NewNatureRemoRepository,
	repository.NewSwitchBotRepository,
	repository.NewOpenWeatherMapRepository,
	wire.Bind(
		new(domain.Clock),
		new(*domain.SystemClock),
	),
)

func AwairRepository(ctx context.Context) domainRepository.AwairRepository {
	wire.Build(binding)
	return nil
}

func NatureRemoRepository(ctx context.Context) domainRepository.NatureRemoRepository {
	wire.Build(binding)
	return nil
}

func SwitchBotRepository(ctx context.Context) domainRepository.SwitchBotRepository {
	wire.Build(binding)
	return nil
}

func OpenWeatherMapRepository(ctx context.Context) domainRepository.OpenWeatherMapRepository {
	wire.Build(binding)
	return nil
}
