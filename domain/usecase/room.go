package usecase

import (
	"context"
	"github.com/y-yu/kindle-clock-go/domain/model"
)

type AllRoomInfo struct {
	AwairRoomInfo      model.AwairRoomInfo
	NatureRemoRoomInfo model.NatureRemoRoomInfo
	SwitchBotMeterInfo model.SwitchBotRoomInfo
	Weather            model.Weather
}

type GetRoomInfoUsecase interface {
	Execute(ctx context.Context) (AllRoomInfo, error)
}
