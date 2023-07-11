package repository

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("github.com/Rikypurnomo/warmup/internal/api/handlers/")
