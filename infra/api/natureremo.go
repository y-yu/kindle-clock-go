package api

import (
	"context"
	"fmt"
	"github.com/go-json-experiment/json"
	"github.com/go-playground/validator/v10"
	"github.com/y-yu/kindle-clock-go/config"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"log/slog"
)

type NatureRemoAPIClientImpl struct {
	config *config.NatureRemoConfiguration
}

func NewNatureRemoAPIClient(c *config.NatureRemoConfiguration) api.NatureRemoAPIClient {
	return &NatureRemoAPIClientImpl{
		config: c,
	}
}

// A is array item type
func parserJsonArray[A any](
	body []byte,
	result *A,
) error {
	var jsonArray []A
	if err := json.Unmarshal(body, &jsonArray); err != nil {
		return err
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	for _, item := range jsonArray {
		if err := validate.Struct(item); err == nil {
			*result = item
			return nil
		}
	}
	slog.Error("item not found in json body", "json array", body)
	return fmt.Errorf("item not found")
}

func (n *NatureRemoAPIClientImpl) GetLatestAllDevicesEvents(ctx context.Context) (api.NatureRemoLatestEvent, error) {
	url := fmt.Sprintf(
		"%s/1/devices",
		n.config.NatureRemoEndpointURL,
	)
	data, err := GetRequestAPI(
		ctx,
		url,
		n.config.OAuthToken,
		func(body []byte, result *api.NatureRemoLatestEvent) error {
			return parserJsonArray(body, result)
		},
	)
	if err != nil {
		return api.NatureRemoLatestEvent{}, err
	}
	return data, nil
}

func (n *NatureRemoAPIClientImpl) GetLatestSmartMeterData(ctx context.Context) (api.NatureRemoSmartMeterResponse, error) {
	url := fmt.Sprintf(
		"%s/1/appliances",
		n.config.NatureRemoEndpointURL,
	)
	data, err := GetRequestAPI(
		ctx,
		url,
		n.config.OAuthToken,
		func(body []byte, result *api.NatureRemoSmartMeterResponse) error {
			return parserJsonArray(body, result)
		},
	)
	if err != nil {
		return api.NatureRemoSmartMeterResponse{}, err
	}
	return data, nil
}
