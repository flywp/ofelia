package core

import (
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/lmittmann/tint"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(w io.Writer) Logger {
	var logLevel slog.Level
	l := os.Getenv("LOG_LEVEL")
	if err := logLevel.UnmarshalText([]byte(l)); err != nil {
		logLevel = slog.LevelInfo
	}

	return &SlogLogger{
		logger: slog.New(tint.NewHandler(w, &tint.Options{
			Level:     logLevel,
			AddSource: false,
			NoColor: func() bool {
				res, _ := strconv.ParseBool(os.Getenv("LOG_NO_COLOR"))
				return res
			}(),
		})),
	}

}

func (l *SlogLogger) Debug(str string, args ...any) {
	l.logger.Debug(str, args...)
}

func (l *SlogLogger) Error(str string, args ...any) {
	l.logger.Error(str, args...)
}

func (l *SlogLogger) Info(str string, args ...any) {
	l.logger.Info(str, args...)
}

func (l *SlogLogger) Warning(str string, args ...any) {
	l.logger.Warn(str, args...)
}
