package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"robot_demo/bootstrap/config"
	"robot_demo/bootstrap/response"
	"robot_demo/bootstrap/xlogger"
	"robot_demo/http"
	"strconv"
)

/**
 * @Author:lingwang
 * @Description:机器人业务
 * @Version: 1.0.0
 * @Date :2022/2/25 16:31
 */

var Size = int(9)
var Number = int(12)
var Num = int(0)

func RobotService(c *gin.Context) {
	response.SuccessResponse(c, "happy")
}

func ChannelMessage(channel, params string) string {

	return http.HttpPost(config.Config.ROBOT_URL+"/channels/"+channel+"/messages", params)
}

func Shudu() string {
	result := ""
	size := 14
	num := 9
	numArr := make([]int, size)
	arr := make([][]int, num)
	for current := 0; current < num; current++ {
		arr[current] = make([]int, num)
	}
	tag := 0
	for current := 0; current < size; {
		num1, num2, num3 := GetRand(num), GetRand(num), GetRand(num)
		if arr[num1][num2] == 0 && GetOutOf(numArr, num3, size) == true && GetSame(arr, num3, num, num1, num2) == true {
			arr[num1][num2] = num3
			numArr[tag] = num3
			tag++
			current++
		}
	}
	result += "\\n"
	for current := 0; current < num; current++ {
		for tag := 0; tag < num; tag++ {
			if arr[current][tag] == 0 {
				result += "_"
			} else {
				result += fmt.Sprintf("%d", arr[current][tag])
			}
			result += " "
		}
		result += "\\n"
	}
	return result
}

func Story() string {
	return "故事会"
}

func Hahaha() string {
	return "笑话集"
}

func GuessNum() string {
	Num = GetRand(10) + 1
	return "游戏开始，数字已生成"
}

func BornNum() string {
	Num = GetRand(10) + 1
	return "已重新生成数字"
}

func IsNumS(number string) string {
	isNum, err := strconv.Atoi(number)
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return "请输入正确的格式哦：/数字 10"
	}
	if isNum == Num {
		return "恭喜你猜中了"
	} else if isNum > Num {
		return "猜大了一点，请猜小一点哦"
	} else {
		return "猜小了一点，请猜大一点哦"
	}
}
func IsNum(number int64) string {
	if number == int64(Num) {
		return "恭喜你猜中了"
	} else if number > int64(Num) {
		return "猜大了一点，请猜小一点哦"
	} else {
		return "猜小了一点，请猜大一点哦"
	}
}
