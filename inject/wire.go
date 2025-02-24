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

func Initialize(ctx context.Context) domainRepository.AwairRepository {
	wire.Build(
		domain.NewSystemClock,
		api.NewAwairApiClient,
		cache.NewAwairCacheClient,
		repository.NewAwairRepository,
		wire.Bind(
			new(domain.Clock),
			new(*domain.SystemClock),
		),
	)
	return nil
}
