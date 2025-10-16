package base

import (
	"context"
	"log/slog"
	"os"
)

// Logger wrapper autour de slog avec méthodes utilitaires.
type Logger struct {
	*slog.Logger
}

// NewLogger crée une nouvelle instance de Logger.
func NewLogger(level string, format string) *Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				formatted := a.Value.Time().Format("02/01/2006 15:04:05")
				return slog.Attr{
					Key:   "timestamp",
					Value: slog.StringValue(formatted),
				}
			}
			return a
		},
	}

	var handler slog.Handler
	if format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return &Logger{
		Logger: slog.New(handler),
	}
}

// LogError helper pour logger les erreurs avec contexte.
func (l *Logger) LogError(ctx context.Context, operation string, err error, fields ...interface{}) {
	args := []interface{}{
		"operation", operation,
		"error", err.Error(),
	}
	args = append(args, fields...)

	// Ajouter des infos du contexte si disponibles
	if requestID := ctx.Value("request_id"); requestID != nil {
		args = append(args, "request_id", requestID)
	}

	l.ErrorContext(ctx, "Operation failed", args...)
}

// LogInfo helper pour logger les informations.
func (l *Logger) LogInfo(ctx context.Context, operation string, message string, fields ...interface{}) {
	args := []interface{}{
		"operation", operation,
	}
	args = append(args, fields...)

	if requestID := ctx.Value("request_id"); requestID != nil {
		args = append(args, "request_id", requestID)
	}

	l.InfoContext(ctx, message, args...)
}

// LogWarn helper pour logger les avertissements.
func (l *Logger) LogWarn(ctx context.Context, operation string, message string, fields ...interface{}) {
	args := []interface{}{
		"operation", operation,
	}
	args = append(args, fields...)

	if requestID := ctx.Value("request_id"); requestID != nil {
		args = append(args, "request_id", requestID)
	}

	l.WarnContext(ctx, message, args...)
}

// With ajoute un contexte au logger.
func (l *Logger) With(args ...interface{}) *Logger {
	return &Logger{
		Logger: l.Logger.With(args...),
	}
}
