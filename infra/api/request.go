package api

import (
	"context"
	"io"
	"log/slog"
	"net/http"
)

var client = http.DefaultClient

// GetRequestAPI sends GET method HTTP request to the url with oauthToken.
func GetRequestAPI[A any](
	ctx context.Context,
	url string,
	oauthToken string,
	jsonParser func(body []byte, result *A) error,
	headerHandler ...func(req *http.Request),
) (A, error) {
	var result A

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	if len(headerHandler) > 0 {
		for _, h := range headerHandler {
			h(req)
		}
	}
	SetAuthHeader(req, oauthToken)

	resp, err := client.Do(req)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			slog.Error(
				"Request body close error!",
				"err", err,
				"method", http.MethodGet,
				"url", url,
			)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = jsonParser(body, &result)
	if err != nil {
		slog.Error(
			"JSON body parse error!",
			"err", err,
			"body", string(body),
			"method", http.MethodGet,
			"url", url,
		)
		return result, err
	}
	return result, nil
}
