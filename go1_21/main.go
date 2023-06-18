package main

import (
	"log/slog"
	"os"
)

func main() {
	// log/slog
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	l.Info("use json handler", slog.Bool("boolkey", true))
	slog.Info("before SetDefault", slog.Bool("boolkey", true))

	slog.SetDefault(l.With(slog.String("withkey", "withvalue")))

	slog.Warn("warn")
}
