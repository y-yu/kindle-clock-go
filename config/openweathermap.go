package config

type OpenWeatherMapConfiguration struct {
	OpenWeatherMapEndPointURL string `env:"OPEN_WEATHER_MAP_ENDPOINT_URL, default=https://api.openweathermap.org"`
	AppID                     string `env:"OPEN_WEATHER_MAP_APP_ID, required"`
	Lat                       string `env:"OPEN_WEATHER_MAP_LAT, required"`
	Lon                       string `env:"OPEN_WEATHER_MAP_LON, required"`
}
