package web

import (
	"github.com/gin-gonic/gin"
	"robot_demo/service"
)

/**
 * @Author:lingwang
 * @Description:机器人接口
 * @Version: 1.0.0
 * @Date :2022/2/25 16:29
 */

func RobotTest(c *gin.Context) {
	service.RobotService(c)
}
