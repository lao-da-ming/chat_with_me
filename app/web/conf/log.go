package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() *zap.Logger {
	projectRootDir, err := os.Getwd()
	if err != nil {
		panic("获取根目录失败" + err.Error())
	}
	logDir := projectRootDir + "/logs"
	_ = os.MkdirAll(logDir, os.ModePerm)
	// 创建自定义配置
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.ErrorLevel), // 日志级别
		Development: true,                                 // 开发模式
		Encoding:    "console",                            // 或 "json"
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 带颜色的级别
			EncodeTime:     zapcore.ISO8601TimeEncoder,       // ISO8601 时间格式
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // 短格式调用者信息
		},
		OutputPaths:      []string{"stdout", projectRootDir + "/logs/myapp.log"}, // 输出位置
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
