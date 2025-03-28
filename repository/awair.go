package repository

import (
	"context"
	"errors"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"github.com/y-yu/kindle-clock-go/domain/model"
	"github.com/y-yu/kindle-clock-go/domain/model/config"
	"github.com/y-yu/kindle-clock-go/domain/repository"
	"github.com/y-yu/kindle-clock-go/infra/cache/proto"
	"log"
	"time"
)

type AwairRepositoryImpl struct {
	config           config.AwairConfiguration
	awairApiClient   api.AwairApiClient
	awairCacheClient domain.CacheClient[*proto.AwairDataModel]
	clock            domain.Clock
}

func NewAwairRepository(
	ctx context.Context,
	awairApiClient api.AwairApiClient,
	awairCacheClient domain.CacheClient[*proto.AwairDataModel],
	clock domain.Clock,
) repository.AwairRepository {
	var c config.AwairConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	return &AwairRepositoryImpl{
		config:           c,
		awairApiClient:   awairApiClient,
		awairCacheClient: awairCacheClient,
		clock:            clock,
	}
}

const (
	compTemperature = "temp"
	compHumidity    = "humid"
	compCo2         = "co2"
	compVoc         = "voc"
	compPm25        = "pm25"
)

func (a *AwairRepositoryImpl) GetRoomInfo(ctx context.Context) (model.AwairRoomInfo, error) {
	cached, err := a.awairCacheClient.Get(ctx, a.config.CacheKeyName)
	if err != nil {
		return model.AwairRoomInfo{}, err
	}
	if cached != nil {
		info := model.AwairRoomInfo{
			Score:       model.Score(cached.Score),
			Temperature: model.Temperature(cached.Temperature),
			Humidity:    model.Humidity(cached.Humidity),
			Co2:         model.Co2(cached.Co2),
			Voc:         model.Voc(cached.Voc),
			Pm25:        model.Pm25(cached.Pm25),
		}
		return info, nil
	}
	response, err := a.awairApiClient.GetLatestAirData(ctx)
	if err != nil {
		return model.AwairRoomInfo{}, err
	}
	if len(response.Data) == 0 || len(response.Data[0].Sensors) == 0 {
		return model.AwairRoomInfo{}, errors.New("empty response")
	}
	data := response.Data[0]
	result := model.AwairRoomInfo{
		Score: model.Score(data.Score),
	}
	for _, s := range response.Data[0].Sensors {
		if s.Comp == compTemperature {
			result.Temperature = model.Temperature(float32(s.Value.(float64)))
		} else if s.Comp == compHumidity {
			result.Humidity = model.Humidity(float32(s.Value.(float64)))
		} else if s.Comp == compCo2 {
			result.Co2 = model.Co2(int32(s.Value.(float64)))
		} else if s.Comp == compVoc {
			result.Voc = model.Voc(int32(s.Value.(float64)))
		} else if s.Comp == compPm25 {
			result.Pm25 = model.Pm25(float32(s.Value.(float64)))
		}
	}
	nowMilliSeconds := a.clock.Now().UnixNano() / int64(time.Millisecond)
	protoData := proto.AwairDataModel{
		Score:                 int32(result.Score),
		Temperature:           float64(result.Temperature),
		Humidity:              float64(result.Humidity),
		Co2:                   int32(result.Co2),
		Voc:                   int32(result.Voc),
		Pm25:                  float64(result.Pm25),
		CreatedAtMilliseconds: nowMilliSeconds,
	}
	err = a.awairCacheClient.Set(ctx, a.config.CacheKeyName, &protoData, a.config.CacheExpire)
	if err != nil {
		return model.AwairRoomInfo{}, err
	}

	return result, nil
}
