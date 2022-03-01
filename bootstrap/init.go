package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"robot_demo/bootstrap/config"
	"robot_demo/bootstrap/exception"
	"robot_demo/bootstrap/response"
	"robot_demo/bootstrap/router"
	"robot_demo/bootstrap/websocket"
	"robot_demo/bootstrap/xlogger"
	"robot_demo/http"
	"robot_demo/pojo"
)

/**
 * @Author:lingwang
 * @Description:项目初始化配置
 * @Version: 1.0.0
 * @Date :2022/2/25 15:06
 */

//项目初始化
func Init() {

	//配置文件
	config.ConfigInit()

	//日志
	xlogger.LoggerInit()

	//路由
	r := gin.Default()
	router.RouterInit(r)

	//全局异常处理
	r.NoMethod(response.HandleNotFound)
	r.NoRoute(response.HandleNotFound)
	r.Use(exception.ExceptionInit())

	http.HttpGet(config.Config.ROBOT_URL + "/users/@me")
	websocketUrl, err := http.HttpGet(config.Config.ROBOT_URL + "/gateway")
	if err == nil {
		pojo.WebUrl.Url = websocketUrl[8 : len(websocketUrl)-2]
	} else {
		xlogger.ErrorLogger.Error(err.Error())
		os.Exit(1)
	}

	//websocket初始化，分别为botGo的官方应用与底层连接
	//websocket.WebSocketPing()
	websocket.WebsocketInit()

	//设置端口
	port := fmt.Sprintf(":%d", config.Config.PORT)
	r.Run(port)

}
