package logger

import (
	"context"
	"fmt"
	"io"
)

// Logger 是一个记录器接口，提供带级别的日志记录功能。
type Logger interface {
	Trace(v ...any)
	Debug(v ...any)
	Info(v ...any)
	Notice(v ...any)
	Warn(v ...any)
	Error(v ...any)
	Fatal(v ...any)
}

// FormatLogger 是一个日志记录器接口，它以某种格式输出日志。
type FormatLogger interface {
	Tracef(format string, v ...any)
	Debugf(format string, v ...any)
	Infof(format string, v ...any)
	Noticef(format string, v ...any)
	Warnf(format string, v ...any)
	Errorf(format string, v ...any)
	Fatalf(format string, v ...any)
}

// CtxLogger 是一个记录器接口，它接受上下文参数并使用格式输出日志。
type CtxLogger interface {
	CtxTracef(ctx context.Context, format string, v ...any)
	CtxDebugf(ctx context.Context, format string, v ...any)
	CtxInfof(ctx context.Context, format string, v ...any)
	CtxNoticef(ctx context.Context, format string, v ...any)
	CtxWarnf(ctx context.Context, format string, v ...any)
	CtxErrorf(ctx context.Context, format string, v ...any)
	CtxFatalf(ctx context.Context, format string, v ...any)
}

// Control 提供配置记录器的方法。
type Control interface {
	SetLevel(Level)
	SetOutput(io.Writer)
}

// FullLogger 是 Logger， FormatLogger， CtxLogger 和 Control 的组合。
type FullLogger interface {
	Logger
	FormatLogger
	CtxLogger
	Control
}

// Level 定义日志消息的优先级。
// 当为日志记录器配置了级别时，将不会输出具有较低日志级别（通过整数比较较小）的任何日志消息。
type Level int

// 日志级别。
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelFatal
)

var strs = []string{
	"[Trace] ",
	"[Debug] ",
	"[Info] ",
	"[Notice] ",
	"[Warn] ",
	"[Error] ",
	"[Fatal] ",
}

func (lv Level) toString() string {
	if lv >= LevelTrace && lv <= LevelFatal {
		return strs[lv]
	}
	return fmt.Sprintf("[?%d] ", lv)
}
