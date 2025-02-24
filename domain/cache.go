package domain

import (
	"context"
	"time"
)

type CacheClient[A ProtoMarshalUnmarshal] interface {
	Get(ctx context.Context, key string) (A, error)
	Set(ctx context.Context, key string, value A, expiration time.Duration) error
}
