package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.opentelemetry.io/otel/attribute"
	otelBridge "go.opentelemetry.io/otel/bridge/opentracing"
	"go.opentelemetry.io/otel/exporters/jaeger" //nolint:staticcheck // This is deprecated and will be removed in the next release.
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"net/http"
)

const (
	service     = "trace-demo"
	environment = "production"
	id          = 1
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func main() {
	tp, _ := tracerProvider("")
	otelTracer := tp.Tracer("tracer_name")

	aa, _ := otelBridge.NewTracerPair(otelTracer)
	aa.SetTextMapPropagator(propagation.TraceContext{})

	header := http.Header{}
	header.Set("traceparent", "11-38ffc3738102c27003f3a627207cd728-38ffc3738102c270-22")

	span, _ := aa.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	fmt.Println("abcd", span)
	otelTracer.Start(context.Background(), "aa")
}
