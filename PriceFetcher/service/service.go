package service

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcherService is an interface that can fetch a price
type PriceFetcherService interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcherService interface.
type PriceFetcher struct{}

func NewPriceFetcherService() *PriceFetcher {
	return &PriceFetcher{}
}

func (svc *PriceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return mockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64 {
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

func mockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100 * time.Millisecond)
	price, err := priceMocks[ticker]

	if !err {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}