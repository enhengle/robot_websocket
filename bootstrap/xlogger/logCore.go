package xlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
 * @Author:lingwang
 * @Description:日志
 * @Version: 1.0.0
 * @Date :2022/2/25 15:24
 */

func NewLogger(filePath, serviceName string, level zapcore.Level,
	maxSize, maxBackups, maxAge int, compress bool) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

func newCore(filePath string, level zapcore.Level,
	maxSize, maxBackup, maxAge int, compress bool) zapcore.Core {
	hook := lumberjack.Logger{
		Filename:   filePath,  //日志文件路径
		MaxSize:    maxSize,   //最大文件保存大小（M）
		MaxAge:     maxBackup, //日志备份最多数量
		MaxBackups: maxAge,    //最多保存天数
		Compress:   compress,  //是否压缩
	}

	//设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "time",
		TimeKey:        "logger",
		NameKey:        "lineNum",
		CallerKey:      "stacktrace",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, //小写
		EncodeTime:     zapcore.ISO8601TimeEncoder,    //UTF时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, //全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           //编辑器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), //打印到控制台和文件
		atomicLevel, //日志级别
	)
}
