package config

import (
	"io"
	"log/slog"
)

func NewLogger(w io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	}))
}
