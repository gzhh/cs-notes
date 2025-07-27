package main

import (
	"context"
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

type contextKey int

const (
	TraceCtxKey contextKey = iota + 1
)

type MyHandler struct {
	slog.Handler
}

func (h MyHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(TraceCtxKey).(string); ok {
		r.Add("trace_id", slog.StringValue(traceID))
	}

	return h.Handler.Handle(ctx, r)
}

func init() {
	// 创建 lumberjack Logger，配置滚动策略
	logFile := &lumberjack.Logger{
		Filename:   "app.log", // 日志文件名
		MaxSize:    1,         // 每个日志文件最大为 1MB
		MaxBackups: 3,         // 最多保留 3 个旧文件
		MaxAge:     28,        // 保留 28 天
		Compress:   true,      // 是否压缩旧日志
	}

	var handler slog.Handler
	handler = slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		AddSource: true,
	})

	handler = MyHandler{handler}

	slog.SetDefault(slog.New(handler))
}

/*
1.slog: structural log
2.lumberjack: log rolling
*/
func main() {
	requestID := "123"
	ctx := context.WithValue(context.Background(), TraceCtxKey, requestID)

	// slog.Info("Hello")
	// slog.InfoContext(ctx, "Hello")
	for i := 0; i < 100000; i++ {
		slog.LogAttrs(ctx, slog.LevelInfo, "Hello", slog.String("key", "value"))
	}
}
