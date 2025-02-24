package api

import (
	"context"
	"fmt"
	"github.com/go-json-experiment/json"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"github.com/y-yu/kindle-clock-go/domain/model/config"
	"io"
	"log"
	"net/http"
)

type AwairApiClientImpl struct {
	config config.AwairConfiguration
}

func NewAwairApiClient(ctx context.Context) api.AwairApiClient {
	var c config.AwairConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	return &AwairApiClientImpl{
		config: c,
	}
}

func (a *AwairApiClientImpl) GetLatestAirData(ctx context.Context) (api.AwairAirResponse, error) {
	url := fmt.Sprintf(
		"%s/v1/users/self/devices/%s/%s/air-data/latest?fahrenheit=false",
		a.config.AwairEndpointURL,
		a.config.DeviceType,
		a.config.DeviceID,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return api.AwairAirResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+a.config.OAuthToken)

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var airData api.AwairAirResponse
	if err := json.Unmarshal(body, &airData); err != nil {
		log.Fatal(err)
		return api.AwairAirResponse{}, err
	}
	return airData, nil
}
