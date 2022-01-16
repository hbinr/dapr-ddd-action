package logger

import (
	"log"
	"os"

	"github.com/dapr-ddd-action/pkg/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZap 初始化 zap Logger
func InitZap(cfg *conf.Config) (logger *zap.Logger) {
	var (
		l    zapcore.Level
		core zapcore.Core
	)
	writeSyncer := getLogWriter(
		cfg.Filename,   // 日志文件的位置
		cfg.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		cfg.MaxBackups, // 保留旧文件的最大个数
		cfg.MaxAge,     // 保留旧文件的最大天数
	)
	encoder := getEncoder()
	if err := l.UnmarshalText([]byte(cfg.Level)); err != nil {
		log.Fatalf("main: init config zap log error :%+v\n", err)
	}

	if cfg.Mode == "debug" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	// logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	zap.ReplaceGlobals(logger)
	zap.L().Info("init logger success")
	return
}

// getEncoder 设置zap编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 指定日志将写到哪里去，并使用Lumberjack进行日志切割归档
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
