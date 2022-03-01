package response

/**
 * @Author:lingwang
 * @Description:返回结构
 * @Version: 1.0.0
 * @Date :2022/2/25 16:05
 */

type Response struct {
	Code int         `json:"code"` //返回码
	Msg  string      `json:"msg"`  //信息
	Data interface{} `json:"data"` //结果
}

var (
	ServerError = ResponseError("系统异常")
	NotFount    = HandleNotFound
)
