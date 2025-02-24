package repository

import (
	"context"
	"github.com/y-yu/kindle-clock-go/domain/model"
)

type AwairRepository interface {
	GetRoomInfo(ctx context.Context) (model.AwairRoomInfo, error)
}
