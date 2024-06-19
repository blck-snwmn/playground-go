package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/blck-snwmn/playground-go/ctxutil/ctxutil/ctxutil"
	"github.com/blck-snwmn/playground-go/ctxutil/ctxutil/id"
	"github.com/blck-snwmn/playground-go/ctxutil/ctxutil/logger"
)

func main() {
	ctx := context.Background()

	ctx = ctxutil.WithValue(ctx, id.ID("123"))
	ctx = ctxutil.WithValue(ctx, logger.New())
	id := ctxutil.Value[id.ID](ctx)
	fmt.Println(id)
	lg := ctxutil.Value[*slog.Logger](ctx)
	lg.Info("info log", slog.Int("intkey", 12))
}
