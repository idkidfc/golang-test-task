package main

import (
	"context"
	"github.com/idkidfc/golang-test-task/internal/domain"
	"github.com/idkidfc/golang-test-task/internal/infrastructure/batchProcessor"
	"github.com/idkidfc/golang-test-task/internal/service/external"
	"github.com/idkidfc/golang-test-task/internal/usecase"
	"time"
)

func main() {
	service := external.NewExternalService(100, time.Second)
	client := batchProcessor.NewBatchProcessorClient(service)
	processor := usecase.NewProcessBatch(client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	batch := domain.Batch{domain.Item{}, domain.Item{}}
	err := processor.Execute(ctx, batch)
	if err != nil {
	}

}
