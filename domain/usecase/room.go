package usecase

import (
	"github.com/y-yu/kindle-clock-go/domain/model"
)

type ShowKindleImageUsecaseResult struct {
	AwairRoomInfo      model.AwairRoomInfo
	NatureRemoRoomInfo model.NatureRemoRoomInfo
	SwitchBotMeterInfo model.SwitchBotRoomInfo
	Weather            model.Weather
}

type GetRoomInfoUsecase interface {
	Execute() (ShowKindleImageUsecaseResult, error)
}
