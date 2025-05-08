package api

import (
	"context"
	"fmt"
	"github.com/go-json-experiment/json"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/config"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type OpenWeatherMapAPIClientImpl struct {
	config config.OpenWeatherMapConfiguration
}

var _ api.OpenWeatherMapAPIClient = (*OpenWeatherMapAPIClientImpl)(nil)

func NewOpenWeatherMapAPIClient(ctx context.Context) api.OpenWeatherMapAPIClient {
	var c config.OpenWeatherMapConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	return &OpenWeatherMapAPIClientImpl{
		config: c,
	}
}

func (o *OpenWeatherMapAPIClientImpl) GetLatest(_ context.Context) (api.OpenWeatherMapInfo, error) {
	url := fmt.Sprintf(
		"%s/data/2.5/weather?lat=%s&lon=%s&appid=%s",
		o.config.OpenWeatherMapEndPointURL,
		o.config.Lat,
		o.config.Lon,
		o.config.AppID,
	)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return api.OpenWeatherMapInfo{}, err
	}
	resp, err := client.Do(req)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			slog.Error(
				"Request body close error!",
				"err", err,
				"method", "GET",
				"url", url,
			)
		}
	}()
	var result api.OpenWeatherMapInfo

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
