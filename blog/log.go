package blog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ConfigCustom = iota //自定义
	ConfigDev           //开发
	ConfigProd          //生产
)
const (
	DebugLevel  = zap.DebugLevel  // -1
	InfoLevel   = zap.InfoLevel   // 0, default level
	WarnLevel   = zap.WarnLevel   // 1
	ErrorLevel  = zap.ErrorLevel  // 2
	DPanicLevel = zap.DPanicLevel // 3, used in development log
	PanicLevel  = zap.PanicLevel  // 4 // PanicLevel logs a message, then panics
	FatalLevel  = zap.FatalLevel  // 5 // FatalLevel logs a message, then calls os.Exit(1).
)

var (
	Skip       = zap.Skip
	Binary     = zap.Binary
	Bool       = zap.Bool
	Boolp      = zap.Boolp
	ByteString = zap.ByteString
	Float64    = zap.Float64
	Float64p   = zap.Float64p
	Float32    = zap.Float32
	Float32p   = zap.Float32p
	String     = zap.String
	Stringp    = zap.Stringp
	Uint       = zap.Uint
	Uintp      = zap.Uintp
	Uint8      = zap.Uint8
	Uint8p     = zap.Uint8p
	Uint32     = zap.Uint32
	Uint32p    = zap.Uint32p
	Uint64     = zap.Uint64
	Uint64p    = zap.Uint64p
	Int        = zap.Int
	Intp       = zap.Intp
	Int8       = zap.Int8
	Int8p      = zap.Int8p
	Int32      = zap.Int32
	Int32p     = zap.Int32p
	Int64      = zap.Int64
	Int64p     = zap.Int64p
	Durationp  = zap.Durationp
	Any        = zap.Any
)

type Config = zap.Config
type Level = zapcore.Level
type Field = zap.Field

type Logger struct {
	l *zap.Logger
}

var logger *Logger

type Options struct {
	Default int //配置类型,默认自定义
	Config  Config
}

func init() {
	cfg := newDefaultDevConfig()
	logger = new(Logger)
	// 构建日志
	logger.l, _ = cfg.Build(zap.AddCallerSkip(1))
}

// New 创建客户端 //todo:支持输出文件路径及日志轮转
func New(opt *Options) (*Logger, error) {
	var cfg Config
	var err error
	switch opt.Default {
	case ConfigDev:
		cfg = newDefaultDevConfig()
	case ConfigProd:
		cfg = newDefaultProdConfig()
	default:
		cfg = opt.Config
	}
	logger := new(Logger)
	// 构建日志
	logger.l, err = cfg.Build(zap.AddCallerSkip(1))
	return logger, err
}

// Configure 配置
func Configure(opt *Options) error {
	var err error
	logger, err = New(opt)
	return err
}

// 默认生产配置
func newDefaultProdConfig() Config {
	return Config{
		Level:       zap.NewAtomicLevelAt(InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// 默认开发配置
func newDefaultDevConfig() Config {
	return Config{
		Level:       zap.NewAtomicLevelAt(DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// GetLogger 获取日志实例
func GetLogger() *Logger {
	return logger
}
func Debug(msg string, fields ...Field) {
	logger.l.Debug(msg, fields...)
}
func (l *Logger) Debug(msg string, fields ...Field) {
	l.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	logger.l.Info(msg, fields...)
}
func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	logger.l.Warn(msg, fields...)
}
func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger.l.Error(msg, fields...)
}
func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func DPanic(msg string, fields ...Field) {
	logger.l.DPanic(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	logger.l.Panic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	logger.l.Fatal(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func Sync() error {
	return logger.l.Sync()
}
func (l *Logger) Sync() error {
	return l.l.Sync()
}
