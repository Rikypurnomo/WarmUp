package handlers

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("github.com/Rikypurnomo/warmup/internal/api/handlers/")

func starSpan(ctx context.Context, name string) trace.Span {
	span := trace.SpanFromContext(ctx)
	span.SetName(name)
	return span
}
