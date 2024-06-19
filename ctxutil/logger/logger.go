package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdin,
			&slog.HandlerOptions{AddSource: true},
		),
	)
}
