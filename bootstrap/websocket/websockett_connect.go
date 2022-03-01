package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/tencent-connect/botgo/dto"
	"log"
	"os"
	"robot_demo/bootstrap/config"
	"robot_demo/bootstrap/xlogger"
	"robot_demo/pojo"
	"robot_demo/service"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author:lingwang
 * @Description:连接websocket
 * @Version: 1.0.0
 * @Date :2022/2/27 16:45
 */

var wssConn *websocket.Conn
var connBool bool

//websocket初始化-失败即重连
func WebsocketInit() {

	timeout := pojo.Timtout{
		Op: 1,
		D:  1,
	}
	err := Link()

	go func() {

		for {

			_, message, err := wssConn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				xlogger.ErrorLogger.Error("read:" + err.Error())
				Link()
			}
			timeMessage := pojo.Timtout{}
			json.Unmarshal(message, &timeMessage)
			switch timeMessage.Op {
			case 0:
				readyEvent := pojo.ReadyEvent{}
				json.Unmarshal(message, &readyEvent)
				timeout.D = readyEvent.S
				contentArr := strings.Split(readyEvent.D.Content, "/")
				if len(contentArr) >= 2 {
					contentArr[1] = strings.Replace(contentArr[1], " ", "", -1)
					switch contentArr[1] {
					case "生成数字":
						{
							params := "{\"content\":\"<@!" + readyEvent.D.Author.Id + ">" + service.BornNum() + "\"}"
							xlogger.InfoLogger.Info(service.ChannelMessage(readyEvent.D.ChannelId, params))
						}
					case "数独":
						{
							params := "{\"content\":\"<@!" + readyEvent.D.Author.Id + ">" + service.Shudu() + "\"}"
							xlogger.InfoLogger.Info(service.ChannelMessage(readyEvent.D.ChannelId, params))
						}
					case "故事会":
						{
							params := "{\"content\":\"<@!" + readyEvent.D.Author.Id + ">" + service.Story() + "\"}"
							xlogger.InfoLogger.Info(service.ChannelMessage(readyEvent.D.ChannelId, params))
						}
					case "笑话集":
						{
							params := "{\"content\":\"<@!" + readyEvent.D.Author.Id + ">" + service.Hahaha() + "\"}"
							xlogger.InfoLogger.Info(service.ChannelMessage(readyEvent.D.ChannelId, params))
						}
					case "猜数字开始":
						{
							params := "{\"content\":\"<@!" + readyEvent.D.Author.Id + ">" + service.GuessNum() + "\"}"
							xlogger.InfoLogger.Info(service.ChannelMessage(readyEvent.D.ChannelId, params))
						}
					case "指令测试":
						{
							params := "{\"content\":\"<@!" + readyEvent.D.Author.Id + ">指令测试完成\"}"
							xlogger.InfoLogger.Info(service.ChannelMessage(readyEvent.D.ChannelId, params))
						}
					default:
						{
							contentArr[1] = strings.Replace(contentArr[1], " ", "", -1)
							contentArr[1] = strings.Replace(contentArr[1], "数字", "", -1)
							if num, err := strconv.ParseInt(contentArr[1], 10, 0); err == nil {
								params := "{\"content\":\"<@!" + readyEvent.D.Author.Id + ">" + service.IsNum(num) + "\"}"
								xlogger.InfoLogger.Info(service.ChannelMessage(readyEvent.D.ChannelId, params))
							}
						}
					}
				}
			case 10: // 接收到 hello 后需要开始发心跳
			case 11: // 心跳 ack 不需要业务处理
			case 9:
				Link()
			}

			xlogger.InfoLogger.Info("recv:" + string(message))
		}
	}()

	timeTicker := time.NewTicker(time.Duration(5) * time.Second)
	for {
		select {
		//定时发送心跳
		case <-timeTicker.C:
			err = wssConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				xlogger.ErrorLogger.Error(err.Error())
				//log.Printf("ping error: %s\n", err)
				Link()

			}
			mjson, _ := json.Marshal(timeout)

			if err := sendMessage(mjson); err != nil {
				xlogger.ErrorLogger.Error(err.Error())
				Link()

			}
			/*if err := wssConn.WriteMessage(websocket.TextMessage, mjson); err != nil {
				log.Printf("pong error: %s\n", err)
				return
			}*/

		}
	}
}
func sendMessage(data []byte) error {
	if err := wssConn.WriteMessage(websocket.TextMessage, data); err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		return err
	}
	xlogger.InfoLogger.Info("发送成功 " + string(data))
	return nil
}

func Link() error {
	var err error
	wssConn, _, err = websocket.DefaultDialer.Dial("wss://sandbox.api.sgroup.qq.com/websocket", nil)
	if err != nil {
		xlogger.ErrorLogger.Error(err.Error())
		os.Exit(1)
	}

	wssConn.ReadMessage()
	connBool = true
	xlogger.InfoLogger.Info("url :wss://sandbox.api.sgroup.qq.com/websocket ,connected")
	event := &dto.WSPayload{
		Data: &dto.WSIdentityData{
			Token:   config.Config.BOT_APPID + "." + config.Config.BOT_TOKEN,
			Intents: 1 << 30,
			Shard:   []uint32{0, 1},
			Properties: struct {
				Os      string `json:"$os,omitempty"`
				Browser string `json:"$browser,omitempty"`
				Device  string `json:"$device,omitempty"`
			}{},
		},
	}
	event.OPCode = 2
	data, _ := json.Marshal(event)
	sendMessage(data)

	_, message1, err := wssConn.ReadMessage()
	if err != nil {
		return err
	}
	xlogger.InfoLogger.Info("鉴权返回信息:" + string(message1))
	return nil
}
