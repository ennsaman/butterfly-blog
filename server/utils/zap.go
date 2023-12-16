package utils

import (
	"blog-server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/fs"
	"log"
	"os"
	"time"
)

// Logger 日志指针
var Logger *zap.Logger

// InitLogger 初始化日志
func InitLogger() {
	// 判断日志文件目录是否存在
	if exists, _ := IsPathExists(config.Conf.Zap.Directory); !exists {
		log.Println("当前目录中不存在日志文件夹，创建日志文件夹：", config.Conf.Zap.Directory)
		_ = os.Mkdir(config.Conf.Zap.Directory, fs.ModePerm) // 创建日志文件夹，权限为 777
	}
	zapcore.NewCore(getEncoder(), getWriterSyncer(), getLevelPriority())
	log.Println("Logger 初始化成功")
}

// getLevelPriority 获取日志级别
func getLevelPriority() zapcore.LevelEnabler {
	level, err := zapcore.ParseLevel(config.Conf.Zap.Level)
	if err != nil {
		log.Println("日志级别解析失败：", err)
		return zapcore.InfoLevel
	}
	return level
}

// getWriterSyncer 获取日志输出路径
func getWriterSyncer() zapcore.WriteSyncer {
	file, _ := os.Create(config.Conf.Zap.Directory + "/log.log")
	if config.Conf.Zap.LogInConsole {
		// 同时输出日志到文件和控制台
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(file))
	}
	return zapcore.AddSync(file)
}

// getEncoder 获取编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "message",    // 日志消息键
		LevelKey:      "level",      // 日志级别键
		TimeKey:       "time",       // 时间键
		NameKey:       "logger",     // 日志记录器键
		CallerKey:     "caller",     // 日志调用者的键
		StacktraceKey: "stacktrace", // 堆栈跟踪键
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime:    customTimeEncoder,           // 自定义时间编码器
		EncodeCaller:  zapcore.ShortCallerEncoder,  // 短调用者编码器
	}

	if config.Conf.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // JSON 编码器
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 控制台编码器
}

// customTimeEncoder 自定义时间编码器
func customTimeEncoder(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(time.Format(config.Conf.Zap.Prefix + "2006-01-02 15:04:05"))
}
