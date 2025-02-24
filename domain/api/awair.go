package api

import (
	"context"
	"time"
)

type AwairAirResponse struct {
	Data []struct {
		Timestamp time.Time `json:"timestamp"`
		Score     int       `json:"score"`
		Sensors   []struct {
			Comp  string `json:"comp"`
			Value any    `json:"value"`
		} `json:"sensors"`
		Indices []struct {
			Comp  string `json:"comp"`
			Value any    `json:"value"`
		} `json:"indices"`
	} `json:"data"`
}

type AwairApiClient interface {
	GetLatestAirData(ctx context.Context) (AwairAirResponse, error)
}
