package log

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var once sync.Once

type Options struct {
	AppName    string `json:"app_name"`
	Level      string `json:"level"`
	TimeFormat string `json:"time_format"`
}

func NewOptions() *Options {
	return &Options{
		AppName:    "unknown",
		Level:      "DEBUG",
		TimeFormat: time.RFC3339Nano,
	}
}

func Init(o *Options) {
	once.Do(func() {
		zerolog.TimeFieldFormat = o.TimeFormat

		var w zerolog.LevelWriter
		w = zerolog.MultiLevelWriter(os.Stdout)

		log.Logger = zerolog.New(w).With().Str("app_name", o.AppName).Logger()

		switch o.Level {
		case "DEBUG":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "INFO":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "WARN":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "ERROR":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	})
}

type ctxLogKeyType string

var ctxLogKey ctxLogKeyType = "logFields"

// WithLogValues 对传入的ctx加入items确定的key value
func WithLogValues(ctx context.Context, items ...string) context.Context {
	if len(items) == 0 {
		return ctx
	}

	logFields := map[string]string{}
	for k, v := range fromCtxLogItems(ctx) {
		logFields[k] = v
	}
	for i := 0; i+1 < len(items); i += 2 {
		logFields[items[i]] = items[i+1]
	}

	return context.WithValue(ctx, ctxLogKey, logFields)
}

func fromCtxLogItems(ctx context.Context) map[string]string {
	if ctx == nil {
		return map[string]string{}
	}
	raw := ctx.Value(ctxLogKey)
	if raw == nil {
		return map[string]string{}
	}
	return raw.(map[string]string)
}

func WithLogContext(ctx context.Context, e *zerolog.Event) *zerolog.Event {
	e.Timestamp()

	if ctx == context.Background() {
		return e
	}

	logFields := fromCtxLogItems(ctx)
	if len(logFields) == 0 {
		return e
	}

	for k, v := range logFields {
		e = e.Str(k, v)
	}

	return e
}

func Debug(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Debug())
}

func Info(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Info())
}

func Warn(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Warn())
}

func WarnErr(ctx context.Context, err error) *zerolog.Event {
	return Warn(ctx).Err(err)
}

func Error(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Error())
}

func Err(ctx context.Context, err error) *zerolog.Event {
	return WithLogContext(ctx, log.Err(err))
}

func Fatal(ctx context.Context) *zerolog.Event {
	return WithLogContext(ctx, log.Fatal())
}
