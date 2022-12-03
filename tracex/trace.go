package tracex

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/go-leo/otelx/resourcex"
)

type Trace struct {
	tracerProvider    trace.TracerProvider
	textMapPropagator propagation.TextMapPropagator
}

func New(ctx context.Context, opts ...Option) (*Trace, error) {
	o := new(options)
	o.apply(opts...)
	o.init()
	var ep ExporterProvider
	switch {
	case o.HTTPOptions != nil:
		ep = o.HTTPOptions
	case o.GRPCOptions != nil:
		ep = o.GRPCOptions
	case o.JaegerOptions != nil:
		ep = o.JaegerOptions
	case o.ZipkinOptions != nil:
		ep = o.ZipkinOptions
	case o.WriterOptions != nil:
		ep = o.WriterOptions
	default:
		return nil, errors.New("not found a trace ExporterProvider")
	}
	exporter, err := ep.Exporter(ctx)
	if err != nil {
		return nil, err
	}

	var bcOpts []sdktrace.BatchSpanProcessorOption

	tpOpts := []sdktrace.TracerProviderOption{
		sdktrace.WithBatcher(exporter, bcOpts...),
		sdktrace.WithSampler(sdktrace.ParentBased(newSample(o.SampleRate))),
		sdktrace.WithResource(resourcex.NewResource(ctx, o.Service, o.Resources, o.Attributes...)),
	}
	if o.IDGenerator != nil {
		tpOpts = append(tpOpts, sdktrace.WithIDGenerator(o.IDGenerator))
	}
	if o.SpanProcessor != nil {
		tpOpts = append(tpOpts, sdktrace.WithSpanProcessor(o.SpanProcessor))
	}
	if o.RawSpanLimits != nil {
		tpOpts = append(tpOpts, sdktrace.WithRawSpanLimits(*o.RawSpanLimits))
	}
	tracerProvider := sdktrace.NewTracerProvider(tpOpts...)
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	return &Trace{tracerProvider: tracerProvider, textMapPropagator: propagator}, nil
}

func (trace *Trace) TracerProvider() trace.TracerProvider {
	return trace.tracerProvider
}

func (trace *Trace) TextMapPropagator() propagation.TextMapPropagator {
	return trace.textMapPropagator
}

func newSample(samplingRate float64) sdktrace.Sampler {
	var sampler sdktrace.Sampler
	switch {
	case samplingRate >= 1:
		sampler = sdktrace.AlwaysSample()
	case samplingRate <= 0:
		sampler = sdktrace.NeverSample()
	default:
		sampler = sdktrace.TraceIDRatioBased(samplingRate)
	}
	return sampler
}
