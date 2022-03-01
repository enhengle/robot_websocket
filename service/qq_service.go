package service

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"robot_demo/bootstrap/response"
)

/**
 * @Author:lingwang
 * @Description:QQ小程序-业务
 * @Version: 1.0.0
 * @Date :2022/2/27 0:35
 */

func GetNum(c *gin.Context) {
	response.SuccessResponse(c, GetRand(10)+1)
}

func GetShudu(c *gin.Context) {
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
	/*for current := 0; current < num; current++ {
		fmt.Println(arr[current])
	}*/
	response.SuccessResponse(c, arr)
}

func GetRand(num int) int {
	return rand.Intn(num)
}

func GetOutOf(arr []int, num, max int) bool {
	size := 0
	for current := 0; current < len(arr); current++ {
		if arr[current] == num {
			size++
			if size >= max {
				return false
			}
		}
	}
	return true
}

func GetSame(arr [][]int, num, size, left, right int) bool {
	for current := 0; current < size; current++ {
		if arr[left][current] == num || arr[current][right] == num {
			return false
		}
	}
	return true
}

func GetLie(c *gin.Context) {
	if GetRand(2) == 1 {
		response.SuccessResponse(c, "you")
	} else {
		response.SuccessResponse(c, "me")
	}
}

func GetHahaha(c *gin.Context) {
	response.SuccessResponse(c, "笑话集内容")
}

func GetStory(c *gin.Context) {
	response.SuccessResponse(c, "故事话内容")
}
