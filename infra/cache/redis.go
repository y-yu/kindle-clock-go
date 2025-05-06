package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/domain/model/config"
	"log/slog"
	"sync"
)

var (
	redisClient *redis.Client
	mu          sync.Mutex
)

func InitializeRedisClient(ctx context.Context) error {
	mu.Lock()
	if redisClient != nil {
		return nil
	}
	defer mu.Unlock()

	var c config.RedisConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("Failed to initialize redis client", "err", err)
		return err
	}
	slog.Info("Initialized redis client")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}
