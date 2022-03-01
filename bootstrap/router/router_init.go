package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"robot_demo/web"
)

/**
 * @Author:lingwang
 * @Description:路由
 * @Version: 1.0.0
 * @Date :2022/2/25 15:57
 */

//路由初始化
func RouterInit(router *gin.Engine) {
	robot := router.Group("robot")
	robot.POST("/get", web.RobotTest)
	robot.GET("/test", ping)
	qq := router.Group("qq")
	qq.POST("/getNum", web.GetNum)
	qq.POST("/getShudu", web.GetShudu)
	qq.POST("/getLie", web.GetLie)
	qq.POST("/getHahaha", web.GetHahaha)
	qq.POST("/getStory", web.GetStory)
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
