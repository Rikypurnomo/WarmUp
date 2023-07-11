package exporter

import (
	"context"

	"github.com/Rikypurnomo/warmup/pkg/logger"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
)

func NewOTLP(endpoint string) *metric.Exporter {
	ctx := context.Background()
	traceGrpc, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint(endpoint))

	if err != nil {
		logger.Errorf("failed to create trace exporter: %s", err.Error())
		return nil
	}

	return &traceGrpc
}
