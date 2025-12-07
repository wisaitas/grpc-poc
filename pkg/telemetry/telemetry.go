package telemetry

import (
	"context"
	"log"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	otellog "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

type Telemetry struct {
	TracerProvider *sdktrace.TracerProvider
	LoggerProvider *sdklog.LoggerProvider
}

func Init(ctx context.Context, serviceName, otelEndpoint string) (*Telemetry, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, err
	}

	// Initialize Tracer
	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(otelEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Initialize Logger
	logExporter, err := otlploggrpc.New(ctx,
		otlploggrpc.WithEndpoint(otelEndpoint),
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter)),
		sdklog.WithResource(res),
	)

	global.SetLoggerProvider(lp)

	log.Printf("Telemetry initialized for service: %s", serviceName)

	return &Telemetry{
		TracerProvider: tp,
		LoggerProvider: lp,
	}, nil
}

func (t *Telemetry) Shutdown(ctx context.Context) error {
	if t.TracerProvider != nil {
		if err := t.TracerProvider.Shutdown(ctx); err != nil {
			return err
		}
	}
	if t.LoggerProvider != nil {
		if err := t.LoggerProvider.Shutdown(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Logger is a simple wrapper for OpenTelemetry logging
type Logger struct {
	logger otellog.Logger
	name   string
}

func NewLogger(name string) *Logger {
	return &Logger{
		logger: global.GetLoggerProvider().Logger(name),
		name:   name,
	}
}

// getTraceInfo extracts trace_id and span_id from context
func getTraceInfo(ctx context.Context) (traceID, spanID string) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		traceID = span.SpanContext().TraceID().String()
		spanID = span.SpanContext().SpanID().String()
	}
	return
}

func (l *Logger) Info(ctx context.Context, msg string, attrs ...slog.Attr) {
	record := otellog.Record{}
	record.SetBody(otellog.StringValue(msg))
	record.SetSeverity(otellog.SeverityInfo)

	// Add trace context to log
	traceID, spanID := getTraceInfo(ctx)
	if traceID != "" {
		record.AddAttributes(
			otellog.String("trace_id", traceID),
			otellog.String("span_id", spanID),
		)
		attrs = append(attrs, slog.String("trace_id", traceID), slog.String("span_id", spanID))
	}

	for _, attr := range attrs {
		record.AddAttributes(otellog.String(attr.Key, attr.Value.String()))
	}
	l.logger.Emit(ctx, record)
	slog.InfoContext(ctx, msg, attrsToAny(attrs)...)
}

func (l *Logger) Error(ctx context.Context, msg string, attrs ...slog.Attr) {
	record := otellog.Record{}
	record.SetBody(otellog.StringValue(msg))
	record.SetSeverity(otellog.SeverityError)

	// Add trace context to log
	traceID, spanID := getTraceInfo(ctx)
	if traceID != "" {
		record.AddAttributes(
			otellog.String("trace_id", traceID),
			otellog.String("span_id", spanID),
		)
		attrs = append(attrs, slog.String("trace_id", traceID), slog.String("span_id", spanID))
	}

	for _, attr := range attrs {
		record.AddAttributes(otellog.String(attr.Key, attr.Value.String()))
	}
	l.logger.Emit(ctx, record)
	slog.ErrorContext(ctx, msg, attrsToAny(attrs)...)
}

func (l *Logger) Warn(ctx context.Context, msg string, attrs ...slog.Attr) {
	record := otellog.Record{}
	record.SetBody(otellog.StringValue(msg))
	record.SetSeverity(otellog.SeverityWarn)

	// Add trace context to log
	traceID, spanID := getTraceInfo(ctx)
	if traceID != "" {
		record.AddAttributes(
			otellog.String("trace_id", traceID),
			otellog.String("span_id", spanID),
		)
		attrs = append(attrs, slog.String("trace_id", traceID), slog.String("span_id", spanID))
	}

	for _, attr := range attrs {
		record.AddAttributes(otellog.String(attr.Key, attr.Value.String()))
	}
	l.logger.Emit(ctx, record)
	slog.WarnContext(ctx, msg, attrsToAny(attrs)...)
}

func attrsToAny(attrs []slog.Attr) []any {
	result := make([]any, 0, len(attrs)*2)
	for _, attr := range attrs {
		result = append(result, attr.Key, attr.Value.String())
	}
	return result
}
