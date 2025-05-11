package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log/slog"
)

type FontConfiguration struct {
	DosisFontPath  string `env:"DOSIS_FONT_PATH, default=./etc/Dosis.ttf"`
	RobotoSlabPath string `env:"ROBOTO_SLAB_FONT_PATH, default=./etc/RobotoSlab.ttf"`
}

func NewFontConfiguration(ctx context.Context) *FontConfiguration {
	var c FontConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration on NewFontConfiguration", "err", err)
		panic(err)
	}

	return &c
}
