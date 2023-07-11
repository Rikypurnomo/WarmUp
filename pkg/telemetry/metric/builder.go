package metric

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type CloseFunc func(ctx context.Context) error

type meterProviderBuilder struct {
	name                string
	exporter            metric.Exporter
	histogramBoundaries []float64
}

func NewMeterProviderBuilder(name string) *meterProviderBuilder {
	return &meterProviderBuilder{
		name: name,
	}
}

func (b *meterProviderBuilder) SetExporter(exp metric.Exporter) *meterProviderBuilder {
	b.exporter = exp
	return b
}

func (b *meterProviderBuilder) SetHistogramBoundaries(explicitBoundaries []float64) *meterProviderBuilder {
	b.histogramBoundaries = explicitBoundaries
	return b
}

func (b meterProviderBuilder) Build() (*metric.MeterProvider, CloseFunc, error) {
	// Create a new meter provider
	reader := metric.NewPeriodicReader(
		b.exporter,
		metric.WithInterval(time.Second),
	)

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(b.name),
	)

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(reader),
	)

	return meterProvider, func(ctx context.Context) error {
		cxt, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := meterProvider.Shutdown(cxt); err != nil {
			return err
		}
		return nil
	}, nil

}
