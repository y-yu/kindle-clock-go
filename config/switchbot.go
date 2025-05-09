package config

import "time"

type SwitchBotConfiguration struct {
	SwitchBotEndpointURL string        `env:"SWITCH_BOT_ENDPOINT_URL, default=https://api.switch-bot.com"`
	OAuthToken           string        `env:"SWITCH_BOT_OAUTH_TOKEN, required"`
	OAuthSecret          string        `env:"SWITCH_BOT_OAUTH_SECRET, required"`
	CacheExpire          time.Duration `env:"SWITCH_BOT_CACHE_EXPIRE, default=1h"`
	CacheKeyName         string        `env:"SWITCH_BOT_CACHE_KEY_NAME, default=switch_bot_cache"`
	Interval             time.Duration `env:"SWITCH_BOT_INTERVAL, default=5m"`
}
