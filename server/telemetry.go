package server

import (
	"time"

	"github.com/Rikypurnomo/warmup/pkg/logger"
	"github.com/Rikypurnomo/warmup/pkg/telemetry/metric"
	mexporter "github.com/Rikypurnomo/warmup/pkg/telemetry/metric/exporter"
	"github.com/Rikypurnomo/warmup/pkg/telemetry/trace"
	texporter "github.com/Rikypurnomo/warmup/pkg/telemetry/trace/exporter"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func InitGlobalProvider(name, endpoint string) {
	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		logger.Fatalf("failed to start runtime: %v", err)
	}
	meterExporter := mexporter.NewOTLP(endpoint)
	meterProvider, meterProviderCloseFn, err := metric.NewMeterProviderBuilder(name).
		SetExporter(*meterExporter).
		Build()
	if err != nil {
		logger.Fatalf("failed to create meter provider: %v", err)
	}
	server.metricProviderCloseFn = append(server.metricProviderCloseFn, meterProviderCloseFn)
	otel.SetMeterProvider(meterProvider)

	spanExporter := texporter.NewOTLP(endpoint)
	tracerProvider, tracerProviderCloseFn, err := trace.NewTraceProviderBuilder(name).
		SetExporter(spanExporter).
		Build()
	if err != nil {
		logger.Fatalf("failed to create trace provider: %v", err)
	}
	server.traceProviderCloseFn = append(server.traceProviderCloseFn, tracerProviderCloseFn)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tracerProvider)
}
