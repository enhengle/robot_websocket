package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

/**
 * @Author:lingwang
 * @Description:配置文件结构
 * @Version: 1.0.0
 * @Date :2022/2/25 15:12
 */
type ConfigStruct struct {
	OPENAPI_ROBOT string
	PORT          int
	ROBOT_URL     string
	BOT_APPID     string
	BOT_TOKEN     string
	BOT_SECRET    string
}

var Config ConfigStruct

//配置文件初始化
func ConfigInit() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./bootstrap/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	Config.OPENAPI_ROBOT = viper.GetString("OPENAPI_ROBOT")
	Config.PORT = viper.GetInt("PORT")
	Config.ROBOT_URL = viper.GetString("ROBOT_URL")
	Config.BOT_APPID = viper.GetString("BOT_APPID")
	Config.BOT_TOKEN = viper.GetString("BOT_TOKEN")
	Config.BOT_SECRET = viper.GetString("BOT_SECRET")
}
