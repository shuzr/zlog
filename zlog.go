package zlog

import (
	"io"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func makeWriter(filename string) io.Writer {
	iow, _ := rotatelogs.New(
		filename+".%Y%m%d"+".log",
		rotatelogs.WithRotationCount(30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	return iow
}

func init() {

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "ts",
		CallerKey:  "caller",
		NameKey:    "app",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	infoWriter := makeWriter(filepath.Base(os.Args[0]))
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
	)
	SugarLogger = zap.New(core, zap.AddCaller()).Sugar()
}

func Debug(args ...interface{}) {
	SugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	SugarLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	SugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	SugarLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	SugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	SugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	SugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	SugarLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	SugarLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	SugarLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	SugarLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	SugarLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	SugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	SugarLogger.Fatalf(template, args...)
}
