package api

import (
	"github.com/go-json-experiment/json"
	"github.com/stretchr/testify/require"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"testing"
)

func TestOpenWeatherMapAPIClientImplParseJSON(t *testing.T) {
	t.Run("success to prase JSON as OpenWeatherMapInfo", func(t *testing.T) {
		jsonStr := `{
			"coord": {
				"lon": 114.514,
				"lat": 19.19
			},
			"weather": [
				{
					"id": 502,
					"main": "Rain",
					"description": "heavy intensity rain",
					"icon": "10d"
				}
			],
			"base": "stations",
			"main": {
				"temp": 287.47,
				"feels_like": 287.28,
				"temp_min": 286.46,
				"temp_max": 288.12,
				"pressure": 1007,
				"humidity": 89,
				"sea_level": 1007,
				"grnd_level": 1005
			},
			"visibility": 6000,
			"wind": {
				"speed": 5.66,
				"deg": 20
			},
			"rain": {
				"1h": 4.21
			},
			"clouds": {
				"all": 75
			},
			"dt": 188888888,
			"sys": {
				"type": 2,
				"id": 810,
				"country": "JP",
				"sunrise": 17000000,
				"sunset": 170000000
			},
			"timezone": 32400,
			"id": 987654,
			"name": "Tokyo",
			"cod": 200
		}`
		var actual api.OpenWeatherMapInfo
		err := json.Unmarshal([]byte(jsonStr), &actual)

		require.NoError(t, err)
		require.Len(t, actual.Weather, 1)
		require.Equal(t, "10d", actual.Weather[0].Icon)
		require.Equal(t, int64(188888888), actual.Datetime)
	})
}
