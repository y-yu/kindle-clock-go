package api

import "context"

type NatureRemoLatestEvent struct {
	NewestEvents struct {
		Te struct {
			Val       float32 `json:"val"`
			CreatedAt string  `json:"created_at"`
		} `json:"te" validate:"required"`
		Hu struct {
			Val       int    `json:"val"`
			CreatedAt string `json:"created_at"`
		} `json:"hu" validate:"required"`
	} `json:"newest_events"`
}

type NatureRemoSmartMeterResponse struct {
	SmartMeter struct {
		EchonetliteProperties []struct {
			Name      string `json:"name"`
			Epc       int    `json:"epc" validate:"required"`
			Val       string `json:"val" validate:"required"`
			UpdatedAt string `json:"updated_at"`
		} `json:"echonetlite_properties" validate:"required"`
	} `json:"smart_meter" validate:"required"`
}

type NatureRemoAPIClient interface {
	GetLatestAllDevicesEvents(ctx context.Context) (NatureRemoLatestEvent, error)

	GetLatestSmartMeterData(ctx context.Context) (NatureRemoSmartMeterResponse, error)
}
