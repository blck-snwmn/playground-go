package logbench

import (
	"fmt"
	"io"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

