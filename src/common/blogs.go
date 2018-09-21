package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

func init() {

	Logger = logs.NewLogger(1000)
	Logger.EnableFuncCallDepth(true)
	if beego.AppConfig.String("runmode") == "dev" { //控制台日志输出
		Logger.SetLogger(logs.AdapterConsole)
	} else if beego.AppConfig.String("runmode") == "prod" { //文件日志输出
		Logger.SetLogger(logs.AdapterFile, `{"filename":"c:/logs/log_manager.log","level":5,"maxlines":0,"maxsize":0,"daily":true,"maxdays":15}`)
	} else {
		Logger.SetLogger(logs.AdapterFile, `{"filename":"c:/logs/log_manager.log","level":6,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	}
}
