package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

func main() {
	fmt.Println("Starting golang service")

	// Configuração -------------------------------------------------------------
	ctx := context.Background()
	res, err := newResources(ctx)
	if err != nil {
		panic(err)
	}

	shutdownTracerProvider := newTracerProvider(ctx, res)
	defer shutdownTracerProvider()

	shutdownMeterProvider := newMeterProvider(ctx, res)
	defer shutdownMeterProvider()

	// Utilização ---------------------------------------------------------------
	tracer := otel.Tracer("go-service-tracer")
	meter := otel.Meter("go-service-meter")

	cpuGauge, err := meter.Int64Gauge(
		"cpu.usage",
		metric.WithDescription("Cpu usage."),
		metric.WithUnit("%"),
	)
	if err != nil {
		panic(err)
	}

	_, span := tracer.Start(ctx, "Starting golang service span")
	span.End()

	for x := 0; x > -1; x++ {
		cpuGauge.Record(ctx, rand.Int63n(100))

		fmt.Println("Doing stuff in golang", x)
		_, childSpan := tracer.Start(ctx, fmt.Sprintf("Doing stuff in golang span %d", x))
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		childSpan.End()

	}
}
