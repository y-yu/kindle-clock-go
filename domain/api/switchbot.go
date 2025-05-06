package api

import "context"

const SwitchBotDeviceTypeMeterPlus = "MeterPlus"

type SwitchBotDeviceList struct {
	DeviceId   string `json:"deviceId"`
	DeviceType string `json:"deviceType"`
}

type SwitchBotDevicesResponse struct {
	Body struct {
		DeviceList []SwitchBotDeviceList `json:"deviceList" validate:"required"`
	} `json:"body" validate:"required"`
}

type SwitchBotDeviceStatusResponse struct {
	Body struct {
		DeviceId    string  `json:"deviceId"`
		Humidity    int     `json:"humidity"`
		Temperature float32 `json:"temperature"`
	} `json:"body" validate:"required"`
}

type SwitchBotAPIClient interface {
	GetDevices(ctx context.Context) (SwitchBotDevicesResponse, error)

	GetLatestMeterData(ctx context.Context, deviceID string) (SwitchBotDeviceStatusResponse, error)
}
