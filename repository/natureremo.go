package repository

import (
	"context"
	"github.com/y-yu/kindle-clock-go/config"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"github.com/y-yu/kindle-clock-go/domain/model"
	"github.com/y-yu/kindle-clock-go/domain/repository"
	"strconv"
)

// nowElectricEnergyNumber is said in https://developer.nature.global/jp/how-to-calculate-energy-data-from-smart-meter-values
const nowElectricEnergyNumber = 231

type NatureRemoRepositoryImpl struct {
	client api.NatureRemoAPIClient
	config *config.NatureRemoConfiguration
}

func NewNatureRemoRepository(
	client api.NatureRemoAPIClient,
	c *config.NatureRemoConfiguration,
) repository.NatureRemoRepository {
	return &NatureRemoRepositoryImpl{
		client: client,
		config: c,
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
		Humidity:       model.Humidity(event.NewestEvents.Hu.Val),
		ElectricEnergy: model.ElectricEnergy(electricEnergy),
	}, nil
}
