package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log/slog"
)

type NatureRemoConfiguration struct {
	NatureRemoEndpointURL string `env:"NATURE_REMO_ENDPOINT_URL, default=https://api.nature.global"`
	OAuthToken            string `env:"NATURE_REMO_OAUTH_TOKEN, required"`
}

func NewNatureRemoConfiguration(ctx context.Context) *NatureRemoConfiguration {
	var c NatureRemoConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration on NewNatureRemoConfiguration", "err", err)
		panic(err)
	}

	return &c
}
