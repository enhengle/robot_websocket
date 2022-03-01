package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robot_demo/bootstrap/response"
)

/**
 * @Author:lingwang
 * @Description:异常捕捉
 * @Version: 1.0.0
 * @Date :2022/2/25 16:46
 */

func ExceptionInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			var res *response.Response
			if e, ok := err.(*response.Response); ok {
				res = e
			} else if e, ok := err.(error); ok {
				res = response.ResponseError(e.Error())
			} else {
				res = response.ServerError
			}
			c.JSON(http.StatusOK, res)
			return
		}
		c.Next()
	}
}
