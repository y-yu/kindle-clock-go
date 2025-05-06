package repository

import (
	"context"
	"github.com/y-yu/kindle-clock-go/domain/model"
)

type SwitchBotRepository interface {
	GetRoomInfo(ctx context.Context) (model.SwitchBotRoomInfo, error)
}
