package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/infra/cache/proto"
	"log"
	"time"
)

type AwairCacheClientImpl struct {
	client *redis.Client
}

func NewAwairCacheClient(ctx context.Context) domain.CacheClient[*proto.AwairDataModel] {
	if redisClient == nil {
		err := InitializeRedisClient(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Initialized redis client:  %v", redisClient)
	}
	return &AwairCacheClientImpl{redisClient}
}

func (a *AwairCacheClientImpl) Get(ctx context.Context, key string) (*proto.AwairDataModel, error) {
	bytes, err := a.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	data := &proto.AwairDataModel{}
	err = data.ProtoUnmarshal(bytes)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AwairCacheClientImpl) Set(ctx context.Context, key string, value *proto.AwairDataModel, expiration time.Duration) error {
	bytes, err := value.ProtoMarshal()
	if err != nil {
		return err
	}
	err = a.client.Set(ctx, key, bytes, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
