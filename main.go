// 启动文件
// @author MoGuQAQ
// @version 1.0.0

package main

import (
	"TeleBot/core"
	"TeleBot/global"
	"TeleBot/telebot"
)

func main() {
	global.Log = core.InitLogger()
	global.Log.Infof("程序开始运行...")
	global.Log.Infof("注册机器人...")
	telebot.InitTelegramBot()
}
