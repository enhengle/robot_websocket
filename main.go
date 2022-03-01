package main

import (
	"robot_demo/bootstrap"
	"robot_demo/bootstrap/config"
	"robot_demo/bootstrap/xlogger"
)

/**
 * @Author:lingwang
 * @Description:入口
 * @Version: 1.0.0
 * @Date :2022/2/25 15:02
 */

func main() {
	bootstrap.Init()
	xlogger.InfoLogger.Info("start_demo")
	xlogger.InfoLogger.Info(config.Config.OPENAPI_ROBOT)
}
