package repository

import (
	"context"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"github.com/y-yu/kindle-clock-go/domain/model"
	"github.com/y-yu/kindle-clock-go/domain/repository"
	"strconv"
)

// nowElectricEnergyNumber is said in https://developer.nature.global/jp/how-to-calculate-energy-data-from-smart-meter-values
const nowElectricEnergyNumber = 231

type NatureRemoRepositoryImpl struct {
	client api.NatureRemoApiClient
}

func NewNatureRemoRepository(client api.NatureRemoApiClient) repository.NatureRemoRepository {
	return &NatureRemoRepositoryImpl{
		client: client,
	}
}

func (n *NatureRemoRepositoryImpl) GetRoomInfo(ctx context.Context) (model.NatureRemoRoomInfo, error) {
	event, err := n.client.GetLatestAllDevicesEvents(ctx)
	if err != nil {
		return model.NatureRemoRoomInfo{}, err
	}
	data, err := n.client.GetLatestSmartMeterData(ctx)
	if err != nil {
		return model.NatureRemoRoomInfo{}, err
	}

	electricEnergy := 0
	for _, property := range data.SmartMeter.EchonetliteProperties {
		if property.Epc == nowElectricEnergyNumber {
			electricEnergy, err = strconv.Atoi(property.Val)
			if err != nil {
				return model.NatureRemoRoomInfo{}, err
			}
			break
		}
	}

	return model.NatureRemoRoomInfo{
		Temperature:    model.Temperature(event.NewestEvents.Te.Val),
		Humidity:       model.Humidity(event.NewestEvents.Te.Val),
		ElectricEnergy: model.ElectricEnergy(electricEnergy),
	}, nil
}
