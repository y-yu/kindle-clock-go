package api

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/go-json-experiment/json"
	"github.com/google/uuid"
	"github.com/y-yu/kindle-clock-go/config"
	"github.com/y-yu/kindle-clock-go/domain/api"
	"net/http"
	"time"
)

type SwitchBotAPIClientImpl struct {
	config *config.SwitchBotConfiguration
}

func NewSwitchBotAPIClient(c *config.SwitchBotConfiguration) api.SwitchBotAPIClient {
	return &SwitchBotAPIClientImpl{
		config: c,
	}
}

var _ api.SwitchBotAPIClient = (*SwitchBotAPIClientImpl)(nil)

// https://github.com/OpenWonderLabs/SwitchBotAPI?tab=readme-ov-file#authentication]]
func requestWithAuthorization[A any](
	ctx context.Context,
	url string,
	oauthToken string,
	oauthSecret string,
	jsonParser func(body []byte, result *A) error,
) (A, error) {
	var result A

	nonce, err := uuid.NewUUID()
	if err != nil {
		return result, err
	}
	now := fmt.Sprintf("%d", time.Now().UnixMilli())
	data := fmt.Sprintf("%s%s%s", oauthToken, now, nonce.String())

	mac := hmac.New(sha256.New, []byte(oauthSecret))
	mac.Write([]byte(data))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	result, err = GetRequestAPI(ctx, url, oauthToken, jsonParser, func(req *http.Request) {
		req.Header.Set("sign", sign)
		req.Header.Set("nonce", nonce.String())
		req.Header.Set("t", now)
	})

	return result, nil
}

func (s *SwitchBotAPIClientImpl) GetDevices(ctx context.Context) (api.SwitchBotDevicesResponse, error) {
	url := fmt.Sprintf(
		"%s/v1.1/devices",
		s.config.SwitchBotEndpointURL,
	)
	data, err := requestWithAuthorization(
		ctx,
		url,
		s.config.OAuthToken,
		s.config.OAuthSecret,
		func(body []byte, result *api.SwitchBotDevicesResponse) error {
			return json.Unmarshal(body, result)
		},
	)
	if err != nil {
		return api.SwitchBotDevicesResponse{}, err
	}
	return data, nil
}

func (s *SwitchBotAPIClientImpl) GetLatestMeterData(
	ctx context.Context,
	deviceID string,
) (api.SwitchBotDeviceStatusResponse, error) {
	url := fmt.Sprintf(
		"%s/v1.1/devices/%s/status",
		s.config.SwitchBotEndpointURL,
		deviceID,
	)
	data, err := requestWithAuthorization(
		ctx,
		url,
		s.config.OAuthToken,
		s.config.OAuthSecret,
		func(body []byte, result *api.SwitchBotDeviceStatusResponse) error {
			return json.Unmarshal(body, result)
		},
	)
	if err != nil {
		return api.SwitchBotDeviceStatusResponse{}, err
	}
	return data, err
}
