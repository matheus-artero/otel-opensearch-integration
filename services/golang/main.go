package main

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newTraceExporter(ctx context.Context) (trace.SpanExporter, error) {
	conn, err := grpc.NewClient("otel-collector:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
}

func newTraceProvider(ctx context.Context, traceExporter trace.SpanExporter) *trace.TracerProvider {
	r, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("service-golang-traces"),
		),
	)

	if err != nil {
		panic(err)
	}

	bsp := trace.NewBatchSpanProcessor(traceExporter)
	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(r),
		trace.WithSpanProcessor(bsp),
	)

	return traceProvider
}

func main() {
	fmt.Println("Starting golang service")

	// Configuração -------------------------------------------------------------
	ctx := context.Background()
	exp, err := newTraceExporter(ctx)
	if err != nil {
		fmt.Println("Failed get console exporter.")
		panic(err)
	}

	tracerProvider := newTraceProvider(ctx, exp)
	defer func() { _ = tracerProvider.Shutdown(ctx) }()

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Utilização ---------------------------------------------------------------
	tracer := otel.Tracer("go-service")
	_, span := tracer.Start(ctx, "Starting golang service span")
	span.End()

	for x := 0; x > -1; x++ {
		fmt.Println("Fake scanning")
		_, childSpan := tracer.Start(ctx, fmt.Sprintf("Fake scanning span %d", x))
		time.Sleep(5 * time.Second)
		childSpan.End()
	}
}
