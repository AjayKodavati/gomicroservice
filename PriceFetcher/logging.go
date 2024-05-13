package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"PriceFetcher/service"
)

type loggingService struct {
	next *service.PriceFetcher
}

func NewLoggingService(next *service.PriceFetcher) *loggingService {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}