//go:build wireinject
// +build wireinject

package inject

import (
	"context"
	"github.com/google/wire"
	"github.com/y-yu/kindle-clock-go/domain"
	domainRepository "github.com/y-yu/kindle-clock-go/domain/repository"
	domainUsecase "github.com/y-yu/kindle-clock-go/domain/usecase"
	"github.com/y-yu/kindle-clock-go/infra/api"
	"github.com/y-yu/kindle-clock-go/infra/cache"
	"github.com/y-yu/kindle-clock-go/presenter/clock"
	"github.com/y-yu/kindle-clock-go/presenter/room"
	"github.com/y-yu/kindle-clock-go/repository"
	"github.com/y-yu/kindle-clock-go/usecase"
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
	usecase.NewGetRoomInfoUsecase,
	room.NewRoomInfoHandler,
	clock.NewClockHandler,
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

func GetRoomInfoUsecase(ctx context.Context) domainUsecase.GetRoomInfoUsecase {
	wire.Build(binding)
	return nil
}

func RoomInfoHandler(ctx context.Context) *room.RoomInfoHandler {
	wire.Build(binding)
	return nil
}

func ClockHandler(ctx context.Context) *clock.ClockHandler {
	wire.Build(binding)
	return nil
}
