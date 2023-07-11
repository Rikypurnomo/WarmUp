package exporter

import (
	"context"
	"time"

	"github.com/Rikypurnomo/warmup/pkg/logger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

func NewOTLP(endpoint string) *otlptrace.Exporter {
	ctx := context.Background()
	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithTimeout(10*time.Second),
	)
	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		logger.Errorf("failed to create trace exporter: %s", err.Error())
		return nil
	}
	return traceExp
}
