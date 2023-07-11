package services

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var tracer = otel.Tracer("github.com/Rikypurnomo/warmup/internal/api/services")

var meters = otel.Meter("github.com/Rikypurnomo/warmup/internal/api/services")

// ServicessInit is an interface for the services
var orderAdd, _ = meters.Int64Counter("order_add", metric.WithDescription("Number of orders added"))
