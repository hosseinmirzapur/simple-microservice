package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	// we can put our metrics logic here, like if we want to send data to prometheus
	fmt.Println("pushing metrics to prometheus")

	return s.next.FetchPrice(ctx, ticker)
}
