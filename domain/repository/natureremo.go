package repository

import (
	"context"
	"github.com/y-yu/kindle-clock-go/domain/model"
)

type NatureRemoRepository interface {
	GetRoomInfo(ctx context.Context) (model.NatureRemoRoomInfo, error)
}
