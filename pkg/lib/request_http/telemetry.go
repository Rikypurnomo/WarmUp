package request_http

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("github.com/Rikypurnomo/warmup/pkg/lib/request_http")
