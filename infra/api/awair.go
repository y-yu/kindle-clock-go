package api

import (
	"context"
	"fmt"
	"github.com/go-json-experiment/json"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"github.com/y-yu/kindle-clock-go/domain/model/config"
	"log"
)

type AwairAPIClientImpl struct {
	config config.AwairConfiguration
}

func NewAwairAPIClient(ctx context.Context) api.AwairAPIClient {
	var c config.AwairConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	return &AwairAPIClientImpl{
		config: c,
	}
}

func (a *AwairAPIClientImpl) GetLatestAirData(ctx context.Context) (api.AwairAirResponse, error) {
	url := fmt.Sprintf(
		"%s/v1/users/self/devices/%s/%s/air-data/latest?fahrenheit=false",
		a.config.AwairEndpointURL,
		a.config.DeviceType,
		a.config.DeviceID,
	)
	airData, err := GetRequestAPI(
		ctx,
		url, a.config.OAuthToken,
		func(body []byte, result *api.AwairAirResponse) error {
			return json.Unmarshal(body, result)
		},
	)
	if err != nil {
		return api.AwairAirResponse{}, err
	}
	return airData, nil
}
