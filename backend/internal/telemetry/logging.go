package telemetry

import (
	"context"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
)

type LoggingConfig struct {
	Level string `enum:"debug,info,warn,error" default:"info"`
	Type  string `enum:"json,console" default:"console"`
}

// converts lower-case log level string to slog.Level values
// TODO: Explore if the marshall methods in slog.Level could instead do the same
func logLevelFromSting(l string) slog.Level {
	switch l {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelInfo
	case "error":
		return slog.LevelInfo
	default:
		return slog.LevelError
	}
}

type Logger interface {
	Fatal(format string, args ...any)
	Error(format string, args ...any)
	Info(format string, args ...any)
	Debug(format string, args ...any)
	Warn(format string, args ...any)
	With(args ...any) *slog.Logger
}

type SlogLogger struct {
	*slog.Logger
}

// NOTE: Use with caution.
// It's not advisable to use Fatal as its similar to panic
func (l *SlogLogger) Fatal(s string, args ...any) {
	l.Error(s, args...)
	os.Exit(1)
}

func NewLogger(logLevelStr, logType string) *SlogLogger {
	return &SlogLogger{Logger: NewSlogLogger(logLevelStr, logType)}
}

func (l *SlogLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(data))
	for k, v := range data {
		attrs = append(attrs, slog.Any(k, v))
	}

	var lvl slog.Level
	switch level {
	case tracelog.LogLevelTrace:
		lvl = slog.LevelDebug - 1
		attrs = append(attrs, slog.Any("PGX_LOG_LEVEL", level))
	case tracelog.LogLevelDebug:
		lvl = slog.LevelDebug
	case tracelog.LogLevelInfo:
		lvl = slog.LevelInfo
	case tracelog.LogLevelWarn:
		lvl = slog.LevelWarn
	case tracelog.LogLevelError:
		lvl = slog.LevelError
	default:
		lvl = slog.LevelError
		attrs = append(attrs, slog.Any("INVALID_PGX_LOG_LEVEL", level))
	}
	l.LogAttrs(context.Background(), lvl, msg, attrs...)
}

// NOTE: this is only needed when setting the default logger using the
//
//	slog.SetDefault method.
func NewSlogLogger(logLevelStr, logType string) *slog.Logger {
	logLevel := &slog.LevelVar{}
	logOpts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}

	logLevel.Set(logLevelFromSting(logLevelStr))
	if logLevel.Level() == slog.LevelDebug {
		logOpts.AddSource = true
	}

	handler := func() slog.Handler {
		if logType == "json" {
			return slog.NewJSONHandler(os.Stderr, &logOpts)
		}
		// return slog.NewTextHandler(os.Stderr, &logOpts)
		// NOTE: Originally we had NewTextHandler, tint just add colors
		return tint.NewHandler(os.Stderr, &tint.Options{
			AddSource: logOpts.AddSource,
			Level:     logOpts.Level,
		})
	}()
	return slog.New(handler)
}
