package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log/slog"
)

type AuthenticationConfiguration struct {
	Token        string `env:"AUTH_TOKEN"`
	QueryKeyName string `env:"AUTH_QUERY_KEY_NAME, default=auth_token"`
}

func NewAuthenticationConfiguration(ctx context.Context) *AuthenticationConfiguration {
	var c AuthenticationConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration on NewAuthenticationConfiguration", "err", err)
		panic(err)
	}

	return &c
}
