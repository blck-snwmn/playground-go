package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	{
		fmt.Println("----")
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
	{
		fmt.Println("----")
		l := slog.New(NewLogger(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
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
}

var _ slog.Handler = (*ContextLogger)(nil)

type ContextLogger struct {
	h slog.Handler
}

// Enabled implements slog.Handler
func (cl *ContextLogger) Enabled(ctx context.Context, l slog.Level) bool {
	return cl.h.Enabled(ctx, l)
}

// Handle implements slog.Handler
func (cl *ContextLogger) Handle(ctx context.Context, r slog.Record) error {
	return cl.h.Handle(ctx, r)
}

// WithAttrs implements slog.Handler
func (cl *ContextLogger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextLogger{h: cl.h.WithAttrs(attrs)}
}

// WithGroup implements slog.Handler
func (cl *ContextLogger) WithGroup(name string) slog.Handler {
	return &ContextLogger{h: cl.h.WithGroup(name)}
}

func NewLogger(h slog.Handler) *ContextLogger {
	return &ContextLogger{h: h}
}
