package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	// l := slog.NewJSONHandler(os.Stdout)
	l.Info("use json handler", slog.Bool("boolkey", true))
	slog.Info("before SetDefault", slog.Bool("boolkey", true))

	slog.SetDefault(l.With(slog.String("withkey", "withvalue")))

	slog.Info("info log",
		slog.Int("intkey", 12),
		slog.Group("group",
			slog.String("key string", "s1"),
			slog.String("key string", "s2"),
		),
	)
	slog.Error("error")
	slog.Warn("warn")

}
