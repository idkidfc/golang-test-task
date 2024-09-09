package batchProcessor

import (
	"context"
	"github.com/idkidfc/golang-test-task/internal/domain"
	"github.com/idkidfc/golang-test-task/internal/infrastructure/rateLimiter"
	"github.com/idkidfc/golang-test-task/internal/service/external"
)

type Client struct {
	service external.Service
	limiter *rateLimiter.TokenBucket
}

func NewBatchProcessorClient(service external.Service) *Client {
	n, p := service.GetLimits()
	return &Client{
		service: service,
		limiter: rateLimiter.NewTokenBucket(n, p),
	}
}

func (c Client) Process(ctx context.Context, batch domain.Batch) error {
	if err := c.limiter.Wait(ctx); err != nil {
		return err
	}
	return c.service.Process(ctx, batch)
}
