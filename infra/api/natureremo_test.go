package api

import (
	"github.com/stretchr/testify/require"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"testing"
)

func TestParser(t *testing.T) {
	json := `[
  {
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
  }
]`

	t.Run("success to parse", func(t *testing.T) {
		var result api.NatureRemoLatestEvent
		err := parser([]byte(json), &result)
		require.NoError(t, err)
		require.Equal(t, float32(17.6), result.NewestEvents.Te.Val)
		require.Equal(t, 50, result.NewestEvents.Hu.Val)
	})
}
