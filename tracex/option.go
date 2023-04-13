package tracex

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-leo/otelx/resourcex"
)

type options struct {
	// Service 服务
	Service *resourcex.Service
	// Attributes trace需要一些额外的信息
	Attributes []attribute.KeyValue
	// Resources 资源
	Resources int
	// HTTPOptions
	HTTPOptions *HTTPOptions
	// GRPCOptions
	GRPCOptions *GRPCOptions
	// JaegerOptions
	JaegerOptions *JaegerOptions
	// ZipkinOptions
	ZipkinOptions *ZipkinOptions
	// WriterOptions
	WriterOptions *WriterOptions
	// SampleRate 采样率
	SampleRate float64
	// IDGenerator 自定义id生成器
	IDGenerator sdktrace.IDGenerator
	// SpanProcessor 自定义span处理器
	SpanProcessor sdktrace.SpanProcessor
	// RawSpanLimits
	RawSpanLimits *sdktrace.SpanLimits
	Propagators   []propagation.TextMapPropagator
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Propagators == nil {
		o.Propagators = []propagation.TextMapPropagator{
			propagation.Baggage{},
			propagation.TraceContext{},
		}
	}
}

type Option func(o *options)

func Service(svc *resourcex.Service) Option {
	return func(o *options) {
		o.Service = svc
	}
}

func Attributes(attrs ...attribute.KeyValue) Option {
	return func(o *options) {
		o.Attributes = attrs
	}
}

func Resources(res int) Option {
	return func(o *options) {
		o.Resources = res
	}
}

func Jaeger(jaegerOptions *JaegerOptions) Option {
	return func(o *options) {
		o.JaegerOptions = jaegerOptions
	}
}

func Zipkin(zipkinOptions *ZipkinOptions) Option {
	return func(o *options) {
		o.ZipkinOptions = zipkinOptions
	}
}

func Writer(writerOptions *WriterOptions) Option {
	return func(o *options) {
		o.WriterOptions = writerOptions
	}
}

func GRPC(gRPCOptions *GRPCOptions) Option {
	return func(o *options) {
		o.GRPCOptions = gRPCOptions
	}
}

func HTTP(httpOptions *HTTPOptions) Option {
	return func(o *options) {
		o.HTTPOptions = httpOptions
	}
}

func SampleRate(rate float64) Option {
	return func(o *options) {
		o.SampleRate = rate
	}
}

func IDGenerator(idGen sdktrace.IDGenerator) Option {
	return func(o *options) {
		o.IDGenerator = idGen
	}
}

func SpanProcessor(spanProcessor sdktrace.SpanProcessor) Option {
	return func(o *options) {
		o.SpanProcessor = spanProcessor
	}
}

func RawSpanLimits(limits *sdktrace.SpanLimits) Option {
	return func(o *options) {
		o.RawSpanLimits = limits
	}
}

func Propagators(propagators ...propagation.TextMapPropagator) Option {
	return func(o *options) {
		o.Propagators = append(o.Propagators, propagators...)
	}
}
