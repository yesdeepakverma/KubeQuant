package controller

/**
Reference: https://pkg.go.dev/github.com/prometheus/client_golang@v1.23.2/api/prometheus/v1#example-API-QueryRange

*/

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func NewPrometheusClientAddress(address string) (api.Client, error) {
	client, err := api.NewClient(api.Config{
		Address: PrometheusURL,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	return client, nil
}

func NewPrometheusClient() (api.Client, error) {
	return NewPrometheusClientAddress(PrometheusURL)
}

func QueryRange(client api.Client, query string, start time.Time, end time.Time, step time.Duration) (model.Value, error) {

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), PrometheusQueryTimeout+2*time.Second)
	defer cancel()
	r := v1.Range{
		Start: start,
		End:   end,
		Step:  step,
	}
	result, warnings, err := v1api.QueryRange(ctx, query, r, v1.WithTimeout(PrometheusQueryTimeout))
	if err != nil {
		fmt.Printf("Error querying Prometheus:%s - %v\n", client, err)
		return nil, err

	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	return result, nil
}
