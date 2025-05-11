package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log/slog"
	"time"
)

type RedisConfiguration struct {
	URL     string        `env:"REDIS_URL, default=redis://localhost:6379"`
	Timeout time.Duration `env:"REDIS_TIMEOUT, default=60s"`
}

func NewRedisConfiguration(ctx context.Context) *RedisConfiguration {
	var c RedisConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration on NewRedisConfiguration", "err", err)
		panic(err)
	}

	return &c
}
