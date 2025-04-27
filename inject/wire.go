//go:build wireinject
// +build wireinject

package inject

import (
	"context"
	"github.com/google/wire"
	"github.com/y-yu/kindle-clock-go/domain"
	domainApi "github.com/y-yu/kindle-clock-go/domain/api"
	domainRepository "github.com/y-yu/kindle-clock-go/domain/repository"
	"github.com/y-yu/kindle-clock-go/infra/api"
	"github.com/y-yu/kindle-clock-go/infra/cache"
	"github.com/y-yu/kindle-clock-go/repository"
)

var binding = wire.NewSet(
	domain.NewSystemClock,
	api.NewAwairApiClient,
	api.NewNatureRemoApiClient,
	api.NewSwitchBotApiClient,
	cache.NewAwairCacheClient,
	repository.NewAwairRepository,
	repository.NewNatureRemoRepository,
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

func SwitchBotClient(ctx context.Context) domainApi.SwitchBotApiClient {
	wire.Build(binding)
	return nil
}
