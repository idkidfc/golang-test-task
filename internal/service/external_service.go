package service

import (
	"context"
	"github.com/idkidfc/golang-test-task/internal/domain"
	"time"
)

type ExternalService interface {
	GetLimits() (n uint64, p time.Duration)
	Process(ctx context.Context, batch domain.Batch) error
}
