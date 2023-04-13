package tracex

import (
	"context"
	"net/http"
	"runtime"

	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type ZipkinOptions struct {
	URL string
}

func (o *ZipkinOptions) Exporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConnsPerHost = runtime.GOMAXPROCS(0) + 1
	return zipkin.New(o.URL, zipkin.WithClient(&http.Client{Transport: transport}))
}
