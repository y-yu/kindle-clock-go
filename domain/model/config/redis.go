package config

import "time"

type RedisConfiguration struct {
	URL     string        `env:"REDIS_URL, default=localhost:6379"`
	Timeout time.Duration `env:"REDIS_TIMEOUT, default=60s"`
}
