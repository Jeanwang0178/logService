package services

import (
	"logService/src/common"
	"logService/src/models"
)

/**
 * 1、tailf file文件  2、发送 kafka 3、页面建立webSocket连接 4、监听kafka消息队列，推送页面
 */
func LogFileServiceViewFile(chanKey string, filePath string, msgKey string) {

	gm := models.NewGoRoutineManager()
	go gm.TailfFiles(chanKey, filePath, msgKey)

}

func LogFileServiceStopTail(chanName string) {

	gm := models.NewGoRoutineManager()
	err := gm.StopLoopGoroutine(chanName)
	if err != nil {
		common.Logger.Error("gm StopLoopGoroutine failed : %v ", err)
	}

}
