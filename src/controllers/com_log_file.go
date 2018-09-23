package controllers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"logService/src/common"
	"logService/src/services"
)

type LogFileController struct {
	BaseController
}

var upgrader = websocket.Upgrader{}

// @router /startTail [post]
func (ctl *LogFileController) StartTail() {

	bodyMap := make(map[string]interface{})
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &bodyMap)

	if bodyMap["filePath"] == nil || bodyMap["chanName"] == nil || bodyMap["msgKey"] == nil {
		common.Logger.Error("missing param filePath || %v ,chanName || %v ,msgKey || %v ", bodyMap["filePath"], bodyMap["chanName"], bodyMap["msgKey"])
		return
	}
	filePath := bodyMap["filePath"].(string)
	chanName := bodyMap["chanName"].(string)
	msgKey := bodyMap["msgKey"].(string)

	services.LogFileServiceViewFile(chanName, filePath, msgKey)

}

// @router /stopTail [post]
func (ctl *LogFileController) StopTail() {

	bodyMap := make(map[string]interface{})
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &bodyMap)

	if bodyMap["chanName"] != nil {
		chanName := bodyMap["chanName"].(string)
		services.LogFileServiceStopTail(chanName)
	} else {
		common.Logger.Error("missing param chanName ")
	}

}
