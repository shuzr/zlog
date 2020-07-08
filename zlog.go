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

var logger *zap.Logger

func Field(k string, v interface{}) zapcore.Field {
	return zap.Any(k, v)
}

func makeWriter(filename string) io.Writer {
	hook, _ := rotatelogs.New(
		filename+".%Y%m%d"+".log",
		rotatelogs.WithRotationCount(30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	return hook
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
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Debug(template string, fields ...zap.Field) {
	logger.Debug(template, fields...)
}

func Info(template string, fields ...zap.Field) {
	logger.Info(template, fields...)
}

func Warn(template string, fields ...zap.Field) {
	logger.Warn(template, fields...)
}

func Error(template string, fields ...zap.Field) {
	logger.Error(template, fields...)
}

func DPanic(template string, fields ...zap.Field) {
	logger.DPanic(template, fields...)
}

func Panic(template string, fields ...zap.Field) {
	logger.Panic(template, fields...)
}

func Fatal(template string, fields ...zap.Field) {
	logger.Fatal(template, fields...)
}
