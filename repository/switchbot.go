package repository

import (
	"context"
	"errors"
	"github.com/samber/lo"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"github.com/y-yu/kindle-clock-go/domain/model"
	"github.com/y-yu/kindle-clock-go/domain/model/config"
	"github.com/y-yu/kindle-clock-go/domain/repository"
	"github.com/y-yu/kindle-clock-go/infra/cache/proto"
	"log"
)

type SwitchBotRepositoryImpl struct {
	config               config.SwitchBotConfiguration
	switchBotAPIClient   api.SwitchBotAPIClient
	switchBotCacheClient domain.CacheClient[*proto.SwitchBotDevicesDataModel]
	clock                domain.Clock
}

var _ repository.SwitchBotRepository = (*SwitchBotRepositoryImpl)(nil)

func NewSwitchBotRepository(
	ctx context.Context,
	switchBotAPIClient api.SwitchBotAPIClient,
	switchBotCacheClient domain.CacheClient[*proto.SwitchBotDevicesDataModel],
	clock domain.Clock,
) repository.SwitchBotRepository {
	var c config.SwitchBotConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	return &SwitchBotRepositoryImpl{
		config:               c,
		switchBotAPIClient:   switchBotAPIClient,
		switchBotCacheClient: switchBotCacheClient,
		clock:                clock,
	}
}

func (s SwitchBotRepositoryImpl) GetRoomInfo(ctx context.Context) (model.SwitchBotRoomInfo, error) {
	cached, err := s.switchBotCacheClient.Get(ctx, s.config.CacheKeyName)
	if err != nil {
		return model.SwitchBotRoomInfo{}, err
	}

	var deviceId string
	if cached != nil && len(cached.Devices) > 0 {
		device, find := lo.Find(cached.Devices, func(device *proto.SwitchBotDevice) bool {
			return device.DeviceType == api.SwitchBotDeviceTypeMeterPlus
		})
		if !find {
			return model.SwitchBotRoomInfo{}, errors.New("Switch Bot MeterPlus device is not found")
		}
		deviceId = device.DeviceId
	} else {
		devices, err := s.switchBotAPIClient.GetDevices(ctx)
		if err != nil {
			return model.SwitchBotRoomInfo{}, err
		}
		protoDevices := lo.Map(devices.Body.DeviceList, func(device api.SwitchBotDeviceList, _ int) *proto.SwitchBotDevice {
			return &proto.SwitchBotDevice{
				DeviceId:   device.DeviceId,
				DeviceType: device.DeviceType,
			}
		})
		err = s.switchBotCacheClient.Set(
			ctx,
			s.config.CacheKeyName,
			&proto.SwitchBotDevicesDataModel{
				Devices: protoDevices,
			},
			s.config.CacheExpire,
		)
		if err != nil {
			return model.SwitchBotRoomInfo{}, err
		}
		device, find := lo.Find(devices.Body.DeviceList, func(device api.SwitchBotDeviceList) bool {
			return device.DeviceType == api.SwitchBotDeviceTypeMeterPlus
		})
		if !find {
			return model.SwitchBotRoomInfo{}, errors.New("Switch Bot MeterPlus device is not found")
		}
		deviceId = device.DeviceId
	}

	response, err := s.switchBotAPIClient.GetLatestMeterData(ctx, deviceId)
	if err != nil {
		return model.SwitchBotRoomInfo{}, err
	}
	return model.SwitchBotRoomInfo{
		Temperature: model.Temperature(response.Body.Temperature),
		Humidity:    model.Humidity(response.Body.Humidity),
	}, nil
}
