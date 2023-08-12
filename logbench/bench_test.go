package logbench

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/xerrors"
)

func Benchmark_slog(b *testing.B) {
	uid := uuid.NewString()
	l := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true}))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(slog.String("user_id", uid))
		l.Info("this is a test. this is error message",
			slog.String("code", "aaaaaaaa"),
			slog.String("message", "this is error message"),
		)
	}
}

func Benchmark_slog_w_error(b *testing.B) {
	uid := uuid.NewString()
	l := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true}))

	err := do()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(slog.String("user_id", uid))
		l.Info("this is a test. this is error message",
			slog.String("code", "aaaaaaaa"),
			slog.String("error", fmt.Sprintf("%+v", err)),
		)
	}
}

func Benchmark_slog_format(b *testing.B) {
	uid := uuid.NewString()
	l := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true}))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(slog.String("user_id", uid))
		l.Info("this is a test. this is error message",
			slog.String("code", fmt.Sprintf("aa%saa", "aaaa")),
			slog.String("message", "this is error message"),
		)
	}
}

func Benchmark_zap(b *testing.B) {
	uid := uuid.NewString()
	cfg := zap.NewProductionConfig()
	l, _ := zap.NewProduction(
		zap.ErrorOutput(zapcore.AddSync(io.Discard)),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg.EncoderConfig),
				zapcore.AddSync(io.Discard),
				cfg.Level,
			)
		}))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(zap.String("user_id", uid))
		l.Info("this is a test. this is error message",
			zap.String("code", "aaaaaaaa"),
			zap.String("message", "this is error message"),
		)
		l.Sync()
	}
}

func Benchmark_zap_w_error(b *testing.B) {
	uid := uuid.NewString()
	cfg := zap.NewProductionConfig()
	l, _ := zap.NewProduction(
		zap.ErrorOutput(zapcore.AddSync(io.Discard)),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg.EncoderConfig),
				zapcore.AddSync(io.Discard),
				cfg.Level,
			)
		}))

	err := do()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(zap.String("user_id", uid))
		l.Info("this is a test. this is error message",
			zap.String("code", "aaaaaaaa"),
			zap.String("error", fmt.Sprintf("%+v", err)),
		)
		l.Sync()
	}
}

func Benchmark_zap_last_sync(b *testing.B) {
	uid := uuid.NewString()
	cfg := zap.NewProductionConfig()
	l, _ := zap.NewProduction(
		zap.ErrorOutput(zapcore.AddSync(io.Discard)),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg.EncoderConfig),
				zapcore.AddSync(io.Discard),
				cfg.Level,
			)
		}))
	defer l.Sync()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(zap.String("user_id", uid))
		l.Info("this is a test. this is error message",
			zap.String("code", "aaaaaaaa"),
			zap.String("message", "this is error message"),
		)

	}
}

func Benchmark_zap_sugar(b *testing.B) {
	uid := uuid.NewString()
	cfg := zap.NewProductionConfig()
	ll, _ := zap.NewProduction(
		zap.ErrorOutput(zapcore.AddSync(io.Discard)),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg.EncoderConfig),
				zapcore.AddSync(io.Discard),
				cfg.Level,
			)
		}))
	l := ll.Sugar()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(zap.String("user_id", uid))
		l.Info("this is a test. this is error message",
			zap.String("code", "aaaaaaaa"),
			zap.String("message", "this is error message"),
		)
		ll.Sync()
	}
}

func Benchmark_zap_sugar_w_error(b *testing.B) {
	uid := uuid.NewString()
	cfg := zap.NewProductionConfig()
	ll, _ := zap.NewProduction(
		zap.ErrorOutput(zapcore.AddSync(io.Discard)),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg.EncoderConfig),
				zapcore.AddSync(io.Discard),
				cfg.Level,
			)
		}))
	l := ll.Sugar()

	err := do()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(zap.String("user_id", uid))
		l.Info("this is a test. this is error message",
			zap.String("code", "aaaaaaaa"),
			zap.Error(err),
		)
		ll.Sync()
	}
}

func Benchmark_zap_sugar_last_sync(b *testing.B) {
	uid := uuid.NewString()
	cfg := zap.NewProductionConfig()
	ll, _ := zap.NewProduction(
		zap.ErrorOutput(zapcore.AddSync(io.Discard)),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg.EncoderConfig),
				zapcore.AddSync(io.Discard),
				cfg.Level,
			)
		}))
	l := ll.Sugar()
	defer l.Sync()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l := l.With(zap.String("user_id", uid))
		l.Info("this is a test. this is error message",
			zap.String("code", "aaaaaaaa"),
			zap.String("message", "this is error message"),
		)
	}
}

func do() error {
	return ido()
}

func ido() error {
	return xerrors.New("error")
}

func TestShowZap(t *testing.T) {
	l, _ := zap.NewProduction()
	l.Info("zap: this is a test. this is error message")
	// {"level":"info","ts":1691777910.0278144,"caller":"logbench/bench_test.go:121","msg":"this is a test. this is error message"}
	l.Sugar().Error(do())
	// {"level":"error","ts":1691778280.8860576,"caller":"logbench/bench_test.go:158","msg":"error","stacktrace":"github.com/blck-snwmn/playground-go/logbench.TestXxx\n\t/home/snowman/dev/github.com/blck-snwmn/playground-go/logbench/bench_test.go:158\ntesting.tRunner\n\t/home/snowman/sdk/go1.21.0/src/testing/testing.go:1595"}
}

func TestShowSlog(t *testing.T) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	l.Info("slog: this is a test. this is error message")
	// {"time":"2023-08-12T03:17:48.101016059+09:00","level":"INFO","source":{"function":"github.com/blck-snwmn/playground-go/logbench.TestXxx2","file":"/home/snowman/dev/github.com/blck-snwmn/playground-go/logbench/bench_test.go","line":126},"msg":"this is a test. this is error message"}
	l.Info(fmt.Sprintf("%+v", do()))
	// {"time":"2023-08-12T14:19:10.772777641+09:00","level":"INFO","source":{"function":"github.com/blck-snwmn/playground-go/logbench.TestShowSlog","file":"/home/snowman/dev/github.com/blck-snwmn/playground-go/logbench/bench_test.go","line":166},"msg":"error:\n    github.com/blck-snwmn/playground-go/logbench.ido\n        /home/snowman/dev/github.com/blck-snwmn/playground-go/logbench/bench_test.go:151"}
}
