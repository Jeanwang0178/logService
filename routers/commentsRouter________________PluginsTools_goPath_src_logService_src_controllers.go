package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["logService/src/controllers:LogFileController"] = append(beego.GlobalControllerRouter["logService/src/controllers:LogFileController"],
		beego.ControllerComments{
			Method:           "StartTail",
			Router:           `/startTail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["logService/src/controllers:LogFileController"] = append(beego.GlobalControllerRouter["logService/src/controllers:LogFileController"],
		beego.ControllerComments{
			Method:           "StopTail",
			Router:           `/stopTail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
