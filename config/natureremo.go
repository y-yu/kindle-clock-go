package config

type NatureRemoConfiguration struct {
	NatureRemoEndpointURL string `env:"NATURE_REMO_ENDPOINT_URL, default=https://api.nature.global"`
	OAuthToken            string `env:"NATURE_REMO_OAUTH_TOKEN, required"`
}
