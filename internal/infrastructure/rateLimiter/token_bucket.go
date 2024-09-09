package rateLimiter

import (
	"context"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens     uint64
	limit      uint64
	interval   time.Duration
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(limit uint64, interval time.Duration) *TokenBucket {
	return &TokenBucket{
		tokens:     limit,
		limit:      limit,
		interval:   interval,
		lastRefill: time.Now(),
	}
}
func (r *TokenBucket) Wait(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for {
		if r.tokens > 0 {
			r.tokens--
			return nil
		}

		now := time.Now()
		elapsed := now.Sub(r.lastRefill)
		if elapsed >= r.interval {
			r.tokens = r.limit
			r.lastRefill = now
			continue
		}

		select {
		case <-time.After(r.interval - elapsed):
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
