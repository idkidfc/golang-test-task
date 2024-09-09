package external

import (
	"context"
	"github.com/idkidfc/golang-test-task/internal/domain"
	"github.com/idkidfc/golang-test-task/pkg/errors"
	"time"
)

type RealExternalService struct {
	limitN uint64
	limitP time.Duration
}

func NewExternalService(limitN uint64, limitP time.Duration) *RealExternalService {
	return &RealExternalService{
		limitN: limitN,
		limitP: limitP,
	}
}

func (s *RealExternalService) GetLimits() (uint64, time.Duration) {
	return s.limitN, s.limitP
}

func (s *RealExternalService) Process(ctx context.Context, batch domain.Batch) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		time.Sleep(100 * time.Millisecond)

		if uint64(len(batch)) > s.limitN {
			return errors.ErrBlocked
		}

		for _, item := range batch {
			_ = item
		}

		return nil
	}
}
