package routers

import (
	"github.com/astaxie/beego"
	"logService/src/controllers"
	_ "logService/src/models"
)

func init() {

	ns := beego.NewNamespace("/open",
		beego.NSNamespace("/logFile",
			beego.NSInclude(
				&controllers.LogFileController{},
			),
		),
	)

	beego.AddNamespace(ns)

}
