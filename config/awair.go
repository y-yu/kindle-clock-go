package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log/slog"
	"time"
)

type AwairConfiguration struct {
	DeviceID         string        `env:"AWAIR_DEVICE_ID, default=15817"`
	DeviceType       string        `env:"AWAIR_DEVICE_TYPE, default=awair-r2"`
	AwairEndpointURL string        `env:"AWAIR_ENDPOINT_URL, default=https://developer-apis.awair.is"`
	OAuthToken       string        `env:"AWAIR_OAUTH_TOKEN, required"`
	CacheExpire      time.Duration `env:"AWAIR_CACHE_EXPIRE, default=1h"`
	Interval         time.Duration `env:"AWAIR_INTERVAL, default=5m"`
	CacheKeyName     string        `env:"AWAIR_CACHE_KEY_NAME, default=awair_cache"`
}

func NewAwairConfiguration(ctx context.Context) *AwairConfiguration {
	var c AwairConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration on NewAwairConfiguration", "err", err)
		panic(err)
	}

	return &c
}
