package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/infra/cache/proto"
	"time"
)

type SwitchBotCacheClientImpl struct {
	client *redis.Client
}

var _ domain.CacheClient[*proto.SwitchBotDevicesDataModel] = (*SwitchBotCacheClientImpl)(nil)

func NewSwitchBotCacheClient(ctx context.Context) domain.CacheClient[*proto.SwitchBotDevicesDataModel] {
	err := InitializeRedisClient(ctx)
	if err != nil {
		panic(err)
	}
	return &SwitchBotCacheClientImpl{redisClient}
}

func (s *SwitchBotCacheClientImpl) Get(
	ctx context.Context,
	key string,
) (*proto.SwitchBotDevicesDataModel, error) {
	bytes, err := s.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	data := &proto.SwitchBotDevicesDataModel{}
	err = data.ProtoUnmarshal(bytes)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *SwitchBotCacheClientImpl) Set(
	ctx context.Context,
	key string,
	value *proto.SwitchBotDevicesDataModel,
	expiration time.Duration,
) error {
	bytes, err := value.ProtoMarshal()
	if err != nil {
		return err
	}
	err = s.client.Set(ctx, key, bytes, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
