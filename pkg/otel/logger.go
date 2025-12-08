package otel

import (
	"context"
	"log/slog"

	otellog "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/trace"
)

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
