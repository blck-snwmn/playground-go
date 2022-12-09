package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	l := slog.NewJSONHandler(os.Stdout)
	slog.SetDefault(slog.New(l))
	slog.Info("info log",
		slog.Int("intkey", 12),
		slog.Group("group",
			slog.String("key string", "s1"),
			slog.String("key string", "s2"),
		),
	)
}
