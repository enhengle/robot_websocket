package web

import (
	"github.com/gin-gonic/gin"
	"robot_demo/service"
)

/**
 * @Author:lingwang
 * @Description:QQ小程序接口
 * @Version: 1.0.0
 * @Date :2022/2/27 0:35
 */

func GetNum(c *gin.Context) {
	service.GetNum(c)
}

func GetShudu(c *gin.Context) {
	service.GetShudu(c)
}

func GetLie(c *gin.Context) {
	service.GetLie(c)
}

func GetHahaha(c *gin.Context) {
	service.GetHahaha(c)
}

func GetStory(c *gin.Context) {
	service.GetStory(c)
}
