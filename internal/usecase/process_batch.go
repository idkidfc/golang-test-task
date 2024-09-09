package usecase

import (
	"context"
	"github.com/idkidfc/golang-test-task/internal/domain"
)

type ProcessBatch struct {
	client BatchProcessor
}

type BatchProcessor interface {
	Process(ctx context.Context, batch domain.Batch) error
}

func NewProcessBatch(client BatchProcessor) *ProcessBatch {
	return &ProcessBatch{client: client}
}

func (p ProcessBatch) Execute(ctx context.Context, batch domain.Batch) error {
	return p.client.Process(ctx, batch)
}
