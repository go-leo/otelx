package otelx

import (
	_ "go.opentelemetry.io/otel/bridge/opencensus"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	_ "go.opentelemetry.io/otel/exporters/prometheus"
	_ "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	_ "go.opentelemetry.io/otel/metric"
	_ "go.opentelemetry.io/otel/sdk/metric"

	_ "go.opentelemetry.io/otel"
	_ "go.opentelemetry.io/otel/bridge/opentracing"
	_ "go.opentelemetry.io/otel/exporters/jaeger"
	// _ "go.opentelemetry.io/otel/exporters/otlp/internal/retry"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	_ "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	_ "go.opentelemetry.io/otel/exporters/zipkin"
	// _ "go.opentelemetry.io/otel/schema"
	// _ "go.opentelemetry.io/otel/sdk"
	_ "go.opentelemetry.io/otel/trace"
)
