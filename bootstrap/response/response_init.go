package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"robot_demo/bootstrap/xlogger"
)

/**
 * @Author:lingwang
 * @Description:异常处理
 * @Version: 1.0.0
 * @Date :2022/2/25 16:17
 */

func ResponseError(err string) *Response {
	return &Response{
		Code: 5500,
		Msg:  err,
		Data: nil,
	}
}

func HandleNotFound(c *gin.Context) {
	err := Response{
		Code: c.Writer.Status(),
		Msg:  c.Request.RequestURI,
		Data: nil,
	}
	xlogger.ErrorLogger.Error(fmt.Sprintf("%d %s", err.Code, err.Msg))
	fmt.Println(err)
	c.JSON(http.StatusOK, err)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	response := Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	fmt.Println(response)
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, response *Response) {
	fmt.Println(response)
	c.JSON(http.StatusOK, &response)
}
