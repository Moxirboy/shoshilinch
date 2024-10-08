package log

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	DebugF(template string, args ...interface{})
	Info(args ...interface{})
	InfoF(template string, args ...interface{})
	Warn(args ...interface{})
	WarnF(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicF(template string, args ...interface{})
	Panic(args ...interface{})
	PanicF(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	LogArgsToString(args interface{}) string
}

type logger struct {
	level    string
	encoding string
}

// Logger
type apiLogger struct {
	logger      *logger
	sugarLogger *zap.SugaredLogger
}

// NewLogger - App Logger constructor
func NewLogger(level, encoding string) *apiLogger {
	return &apiLogger{logger: &logger{level: level, encoding: encoding}}
}

// For mapping config logger to email_service logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(log *logger) zapcore.Level {
	level, exist := loggerLevelMap[log.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// InitLogger initialize application logger
func (l *apiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.logger)

	logWriter := zapcore.AddSync(os.Stderr)
	encoderCfg := zap.NewProductionEncoderConfig()

	var encoder zapcore.Encoder
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderCfg.NameKey = "name"

	if l.logger.encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

func (l *apiLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *apiLogger) LogArgsToString(args interface{}) string {
	jsonData, err := json.MarshalIndent(args, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}
	return string(jsonData)
}
func (l *apiLogger) DebugF(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *apiLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *apiLogger) InfoF(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *apiLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *apiLogger) WarnF(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *apiLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *apiLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *apiLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *apiLogger) DPanicF(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *apiLogger) PanicF(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *apiLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
