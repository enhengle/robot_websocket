package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"robot_demo/bootstrap/config"
	"robot_demo/bootstrap/xlogger"
)

/**
 * @Author:lingwang
 * @Description:接口请求
 * @Version: 1.0.0
 * @Date :2022/2/26 17:11
 */

//get请求
func HttpGet(url string) (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Bot "+config.Config.BOT_APPID+"."+config.Config.BOT_TOKEN)
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return "", err
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return "", err
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return "", err
	}
	return string(b), nil
}

//发送消息请求
func HttpPost(url, params string) string {
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(params)))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bot "+config.Config.BOT_APPID+"."+config.Config.BOT_TOKEN)
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return ""
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return ""
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return ""
	}
	return string(b)
}
