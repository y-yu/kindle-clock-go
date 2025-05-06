package usecase

import (
	"context"
	"github.com/y-yu/kindle-clock-go/domain/repository"
	"github.com/y-yu/kindle-clock-go/domain/usecase"
	"sync"
)

type GetRoomInfoUsecaseImpl struct {
	natureRemoRepository     repository.NatureRemoRepository
	switchBotRepository      repository.SwitchBotRepository
	awairRepository          repository.AwairRepository
	openWeatherMapRepository repository.OpenWeatherMapRepository
}

var _ usecase.GetRoomInfoUsecase = (*GetRoomInfoUsecaseImpl)(nil)

func NewGetRoomInfoUsecase(
	natureRemoRepository repository.NatureRemoRepository,
	switchBotRepository repository.SwitchBotRepository,
	awairRepository repository.AwairRepository,
	openWeatherMapRepository repository.OpenWeatherMapRepository,
) usecase.GetRoomInfoUsecase {
	return &GetRoomInfoUsecaseImpl{
		natureRemoRepository:     natureRemoRepository,
		switchBotRepository:      switchBotRepository,
		awairRepository:          awairRepository,
		openWeatherMapRepository: openWeatherMapRepository,
	}
}

func (g *GetRoomInfoUsecaseImpl) Execute(ctx context.Context) (result usecase.AllRoomInfo, err error) {
	var wg sync.WaitGroup

	errChan := make(chan error, 4)
	defer close(errChan)

	wg.Add(1)
	go func() {
		defer wg.Done()

		natureRemoInfo, err := g.natureRemoRepository.GetRoomInfo(ctx)
		if err != nil {
			errChan <- err
		}
		result.NatureRemoRoomInfo = natureRemoInfo
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		switchBotInfo, err := g.switchBotRepository.GetRoomInfo(ctx)
		if err != nil {
			errChan <- err
		}
		result.SwitchBotMeterInfo = switchBotInfo
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		awairInfo, err := g.awairRepository.GetRoomInfo(ctx)
		if err != nil {
			errChan <- err
		}
		result.AwairRoomInfo = awairInfo
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		weather, err := g.openWeatherMapRepository.GetCurrentWeather(ctx)
		if err != nil {
			errChan <- err
		}
		result.Weather = weather
	}()

	wg.Wait()

	select {
	case err = <-errChan:
		return result, err
	default:
		return result, nil
	}
}
