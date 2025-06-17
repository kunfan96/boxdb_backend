package utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// 定义日志时间戳格式
func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

// 寻找一个文件，如果没有则创造文件
func CreateFileIfNotExist(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return os.Create(filePath)
	}

	// 以追加读写模式打开文件
	return os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0644)
}

// 仅允许 INFO 级别日志写入 info.log
func infoLevelEnabler(level zapcore.Level) bool {
	return level == zapcore.InfoLevel
}

func init() {
	logOutputConfig := GetBootstrapConfig().Log

	// create log instance
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		MessageKey:    "message",
		StacktraceKey: "trace",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			// add log info prefix
			encodeTimeLayout(t, "---\n2006-01-02 15:04:05.000", enc)
		},
		EncodeLevel: zapcore.CapitalLevelEncoder,
	})

	// 创建不同级别的日志文件
	infoFile, _ := CreateFileIfNotExist(logOutputConfig.InfoOutput)
	errorFile, _ := CreateFileIfNotExist(logOutputConfig.ErrorOutput)

	infoWriter := zapcore.AddSync(infoFile)
	errorWriter := zapcore.AddSync(errorFile)

	// 创建不同级别的 Core
	infoCore := zapcore.NewCore(encoder, infoWriter, zap.LevelEnablerFunc(infoLevelEnabler))
	errorCore := zapcore.NewCore(encoder, errorWriter, zapcore.ErrorLevel)

	// 使用 Tee 组合多个 Core
	Logger = zap.New(zapcore.NewTee(infoCore, errorCore), zap.AddStacktrace(zapcore.ErrorLevel))

	defer Logger.Sync()
}
