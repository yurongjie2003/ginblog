package Log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

//go get go.uber.org/zap
//go get github.com/gin-contrib/zap
//go get gopkg.in/natefinch/lumberjack.v2

var Logger *zap.Logger

func InitLogger() error {
	// 配置 lumberjack 日志切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./ginblog.log", // 日志文件路径
		MaxSize:    100,             // 每个日志文件的最大大小（MB）
		MaxBackups: 3,               // 保留的旧日志文件最大数量
		MaxAge:     30,              // 保留旧日志文件的最大天数
		Compress:   true,            // 是否压缩旧日志文件
	}

	// 创建 zapcore.WriteSyncer
	writeSyncer := zapcore.AddSync(lumberjackLogger)

	// 配置 zap 编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式

	// 创建 zapcore.Core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 使用 JSON 编码器
		zapcore.NewMultiWriteSyncer( // 多路输出（控制台 + 文件）
			zapcore.AddSync(writeSyncer),
			zapcore.AddSync(zapcore.Lock(os.Stdout)), // 同时输出到控制台
		),
		zapcore.DebugLevel, // 日志级别
	)

	// 创建 Logger
	Logger = zap.New(core, zap.AddCaller()) // 添加调用者信息
	zap.ReplaceGlobals(Logger)              // 替换全局 Logger

	return nil
}
