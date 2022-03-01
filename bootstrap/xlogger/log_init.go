package xlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
 * @Author:lingwang
 * @Description: 初始化日志配置
 * @Version: 1.0.0
 * @Date :2022/2/25 15:40
 */

var InfoLogger *zap.Logger  //常规日志
var ErrorLogger *zap.Logger //错误日志

//日志文件初始化
func LoggerInit() {
	InfoLogger = NewLogger("./log/info.log", "info", zapcore.InfoLevel, 128, 20, 7, true)
	ErrorLogger = NewLogger("./log/error.log", "error", zapcore.InfoLevel, 128, 20, 7, true)
}
