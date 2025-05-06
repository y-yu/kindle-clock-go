package config

type AuthenticationConfiguration struct {
	Token        string `env:"AUTH_TOKEN"`
	QueryKeyName string `env:"AUTH_QUERY_KEY_NAME, default=auth_token"`
}
