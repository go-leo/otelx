package metricx

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"

	"github.com/go-leo/otelx/resourcex"
)

type Metric struct {
	meterProvider *sdkmetric.MeterProvider
}

func NewMetric(ctx context.Context, opts ...Option) (*Metric, error) {
	o := new(options)
	o.apply(opts...)
	o.init()
	var ep ExporterProvider
	switch {
	case o.PrometheusOptions != nil:
		ep = o.PrometheusOptions
	case o.GRPCOptions != nil:
		ep = o.GRPCOptions
	case o.HTTPOptions != nil:
		ep = o.HTTPOptions
	case o.WriterOptions != nil:
		ep = o.WriterOptions
	default:
		return nil, errors.New("not found a metric ExporterProvider")
	}
	exporter, err := ep.Exporter(ctx)
	if err != nil {
		return nil, err
	}
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resourcex.NewResource(ctx, o.Service, o.Resources, o.Attributes...)),
		sdkmetric.WithReader(exporter),
		sdkmetric.WithView(newViews()...),
	)
	return &Metric{meterProvider: meterProvider}, nil
}

func newViews(viewOpts ...ViewOption) []sdkmetric.View {
	var views []sdkmetric.View
	for _, opt := range viewOpts {
		views = append(views, sdkmetric.NewView(opt.Criteria, opt.Mask))
	}
	return views
}

func (metric *Metric) MeterProvider() metric.MeterProvider {
	return metric.meterProvider
}
