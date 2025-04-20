package api

import (
	"github.com/stretchr/testify/require"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("success to parse as NatureRemoLatestEvent", func(t *testing.T) {
		json := `[{
			"name": "Remo mini",
			"id": "e41eafd3-5b33-412e-8dc9-866a6a0b2d6a",
			"created_at": "2021-11-08T14:24:22Z",
			"updated_at": "2024-09-01T00:46:04Z",
			"mac_address": "40:f5:20:97:6c:20",
			"bt_mac_address": "40:f5:20:97:6c:22",
			"serial_number": "2W221050025872",
			"firmware_version": "Remo-mini/1.14.6",
			"temperature_offset": 0,
			"humidity_offset": 0,
			"users": [
				{
					"id": "eb5d79b1-cf00-4290-a0d0-716c73c3bcd3",
					"nickname": "吉村 優",
					"superuser": true
				}
			],
			"newest_events": {
				"te": {
					"val": 18.1,
					"created_at": "2025-03-20T09:17:58Z"
				}
			},
			"online": true
		},
		{
			"name": "Remo 3rd gen",
			"id": "4aacc1ff-cf99-408d-99ea-1cf9936a0e15",
			"created_at": "2020-07-22T03:09:04Z",
			"updated_at": "2024-09-06T03:42:24Z",
			"mac_address": "98:f4:ab:1f:33:d0",
			"bt_mac_address": "98:f4:ab:1f:33:d2",
			"serial_number": "1W320050000152",
			"firmware_version": "Remo/1.14.6",
			"temperature_offset": 0,
			"humidity_offset": 0,
			"users": [
				{
					"id": "eb5d79b1-cf00-4290-a0d0-716c73c3bcd3",
					"nickname": "吉村 優",
					"superuser": true
				}
			],
			"newest_events": {
				"hu": {
					"val": 50,
					"created_at": "2025-03-20T09:26:24Z"
				},
				"il": {
					"val": 21,
					"created_at": "2025-03-20T09:07:06Z"
				},
				"mo": {
					"val": 1,
					"created_at": "2025-01-24T11:22:51Z"
				},
				"te": {
					"val": 17.6,
					"created_at": "2025-03-20T09:00:22Z"
				}
			},
			"online": true
		},
		{
			"name": "Remo E lite",
			"id": "f2b66b45-afdc-427d-aa54-9ae5d9977a84",
			"created_at": "2020-05-14T06:25:48Z",
			"updated_at": "2024-05-19T12:22:54Z",
			"mac_address": "98:f4:ab:1f:73:0c",
			"bt_mac_address": "98:f4:ab:1f:73:0e",
			"serial_number": "4W120040002071",
			"firmware_version": "Remo-E-lite/1.10.0",
			"temperature_offset": 0,
			"humidity_offset": 0,
			"users": [
				{
					"id": "eb5d79b1-cf00-4290-a0d0-716c73c3bcd3",
					"nickname": "吉村 優",
					"superuser": true
				}
			],
			"newest_events": {},
			"online": true
		}]`
		var result api.NatureRemoLatestEvent
		err := parserJsonArray([]byte(json), &result)
		require.NoError(t, err)
		require.Equal(t, float32(17.6), result.NewestEvents.Te.Val)
		require.Equal(t, 50, result.NewestEvents.Hu.Val)
	})

	t.Run("success to parse as NatureRemoSmartMeterResponse", func(t *testing.T) {
		json := `[{
			"id": "2cdfbf2b-4a6e-4284-8091-5f809dacf3ee",
			"device": {
				"name": "Remo 3rd gen",
				"id": "4aacc1ff-cf99-408d-99ea-1cf9936a0e15",
				"created_at": "2020-07-22T03:09:04Z",
				"updated_at": "2024-09-06T03:42:24Z",
				"mac_address": "98:f4:ab:1f:33:d0",
				"bt_mac_address": "98:f4:ab:1f:33:d2",
				"serial_number": "1W320050000152",
				"firmware_version": "Remo/1.14.6",
				"temperature_offset": 0,
				"humidity_offset": 0
			},
			"model": null,
			"type": "IR",
			"nickname": "サーキュレーター",
			"image": "ico_fan",
			"settings": null,
			"aircon": null,
			"signals": [
				{
					"id": "b7b34711-2953-408b-a4de-8870be26c174",
					"name": "オン",
					"image": "ico_on"
				},
				{
					"id": "0ab57d8e-2519-48af-a599-eabed0a8455c",
					"name": "オフ",
					"image": "ico_off"
				},
				{
					"id": "93c93032-8746-4566-92d2-20e245cb3fe6",
					"name": "首ふり",
					"image": "ico_return"
				},
				{
					"id": "4792927c-f356-4a6d-bb7d-af55edbefafa",
					"name": "1",
					"image": "ico_number_1"
				},
				{
					"id": "65a22a87-8ebe-4b94-9032-b4284175f8b6",
					"name": "2",
					"image": "ico_number_2"
				},
				{
					"id": "556b0a02-e2aa-48eb-9689-5c77250ee3ac",
					"name": "強",
					"image": "ico_number_3"
				}
			]
		},
		{
			"id": "197c8f1f-729a-4927-82e0-f4c93c56f99f",
			"device": {
				"name": "Remo E lite",
				"id": "f2b66b45-afdc-427d-aa54-9ae5d9977a84",
				"created_at": "2020-05-14T06:25:48Z",
				"updated_at": "2024-05-19T12:22:54Z",
				"mac_address": "98:f4:ab:1f:73:0c",
				"bt_mac_address": "98:f4:ab:1f:73:0e",
				"serial_number": "4W120040002071",
				"firmware_version": "Remo-E-lite/1.10.0",
				"temperature_offset": 0,
				"humidity_offset": 0
			},
			"model": {
				"id": "7f3de26b-0afa-44fe-8680-7cf67f8bd415",
				"manufacturer": "",
				"name": "Smart Meter",
				"image": "ico_smartmeter"
			},
			"type": "EL_SMART_METER",
			"nickname": "スマートメーター",
			"image": "ico_smartmeter",
			"settings": null,
			"aircon": null,
			"signals": [],
			"smart_meter": {
				"echonetlite_properties": [
					{
						"name": "coefficient",
						"epc": 211,
						"val": "1",
						"updated_at": "2025-04-20T03:40:42Z"
					},
					{
						"name": "cumulative_electric_energy_effective_digits",
						"epc": 215,
						"val": "6",
						"updated_at": "2025-04-20T03:40:42Z"
					},
					{
						"name": "normal_direction_cumulative_electric_energy",
						"epc": 224,
						"val": "98418",
						"updated_at": "2025-04-20T03:40:42Z"
					},
					{
						"name": "cumulative_electric_energy_unit",
						"epc": 225,
						"val": "1",
						"updated_at": "2025-04-20T03:40:42Z"
					},
					{
						"name": "reverse_direction_cumulative_electric_energy",
						"epc": 227,
						"val": "8",
						"updated_at": "2025-04-20T03:40:42Z"
					},
					{
						"name": "measured_instantaneous",
						"epc": 231,
						"val": "215",
						"updated_at": "2025-04-20T03:40:43Z"
					}
				]
			}
		}]`
		var result api.NatureRemoSmartMeterResponse
		err := parserJsonArray([]byte(json), &result)
		require.NoError(t, err)
		require.Len(t, result.SmartMeter.EchonetliteProperties, 6)
		require.Equal(t, result.SmartMeter.EchonetliteProperties[5].Epc, 231)
		require.Equal(t, result.SmartMeter.EchonetliteProperties[5].Val, "215")
	})
}
