package api

import (
	"github.com/go-json-experiment/json"
	"github.com/stretchr/testify/require"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"testing"
)

func TestSwitchBotAPIClientImpl_GetDevices(t *testing.T) {
	t.Run("success to parse JSON as SwitchBotDevicesResponse", func(t *testing.T) {
		jsonStr := `{
			"statusCode": 100,
			"body": {
				"deviceList": [
					{
						"deviceId": "DF1875E81CFD",
						"deviceName": "温湿度計プラス",
						"deviceType": "MeterPlus",
						"enableCloudService": true,
						"hubDeviceId": "F1A5D0A81EBE"
					},
					{
						"deviceId": "DF6CE0D8E4CB",
						"deviceName": "エントランスドア1",
						"deviceType": "Bot",
						"enableCloudService": true,
						"hubDeviceId": "F1A5D0A81EBE"
					},
					{
						"deviceId": "E4469F852F0A",
						"deviceName": "寝室カーテン（右）",
						"deviceType": "Curtain",
						"enableCloudService": true,
						"hubDeviceId": "F1A5D0A81EBE",
						"curtainDevicesIds": [
							"FB9E9257CB47",
							"E4469F852F0A"
						],
						"calibrate": true,
						"group": true,
						"master": false,
						"openDirection": "right"
					},
					{
						"deviceId": "E4AFE0778592",
						"deviceName": "エントランスドア2",
						"deviceType": "Bot",
						"enableCloudService": true,
						"hubDeviceId": "F1A5D0A81EBE"
					},
					{
						"deviceId": "F1A5D0A81EBE",
						"deviceName": "Hub Plus BE",
						"deviceType": "Hub Plus",
						"hubDeviceId": "000000000000"
					},
					{
						"deviceId": "F50A8CDADAEC",
						"deviceName": "Windowsサーバー",
						"deviceType": "Bot",
						"enableCloudService": true,
						"hubDeviceId": "F1A5D0A81EBE"
					},
					{
						"deviceId": "FB9E9257CB47",
						"deviceName": "寝室カーテン",
						"deviceType": "Curtain",
						"enableCloudService": true,
						"hubDeviceId": "F1A5D0A81EBE",
						"curtainDevicesIds": [
							"FB9E9257CB47",
							"E4469F852F0A"
						],
						"calibrate": true,
						"group": true,
						"master": true,
						"openDirection": "left"
					}
				],
				"infraredRemoteList": []
			},
			"message": "success"
		}`

		var actual api.SwitchBotDevicesResponse
		err := json.Unmarshal([]byte(jsonStr), &actual)
		require.NoError(t, err)
	})
}
