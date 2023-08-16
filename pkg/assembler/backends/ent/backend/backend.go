package backend

import (
	"context"
	"fmt"

	"github.com/guacsec/guac/pkg/assembler/backends"
	"github.com/guacsec/guac/pkg/assembler/backends/ent"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slices"

	// Import regular postgres driver
	_ "github.com/lib/pq"
)

var (
	PathContains = slices.Contains[string]
	Errorf       = gqlerror.Errorf
)

// MaxPageSize is the maximum number of results that will be returned in a single query.
const MaxPageSize = 1000

type EntBackend struct {
	backends.Backend
	client *ent.Client
}

var tracer trace.Tracer

func newExporter(ctx context.Context) /* (someExporter.Exporter, error) */ {
	// Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("EntBackend"),
			attribute.String("environment", "production"),
		)),
	)
	return tp, nil
}

func GetBackend(args backends.BackendArgs) (backends.Backend, error) {
	ctx := context.Background()

	// Create a new tracer provider with a batch span processor and the given exporter.
	tp, err := tracerProvider("http://jaeger:14268/api/traces")
	if err != nil {
		return nil, err
	}

	// Handle shutdown properly so nothing leaks.
	// defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// Finally, set the tracer that can be used for this package.
	tracer = tp.Tracer("EntBackend")
	ctx, span := tracer.Start(ctx, "Init Backend")
	defer span.End()

	be := &EntBackend{}
	if args == nil {
		return nil, fmt.Errorf("invalid args: WithClient is required, got nil")
	}

	if client, ok := args.(*ent.Client); ok {
		err := client.Ping(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to ping db: %w", err)
		}

		be.client = client
	} else {
		return nil, fmt.Errorf("invalid args type")
	}

	return be, nil
}
