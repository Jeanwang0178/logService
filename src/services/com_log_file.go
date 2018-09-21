package services

import (
	"logService/src/common"
	"logService/src/models"
)

/**
 * 1、tailf file文件  2、发送 kafka 3、页面建立webSocket连接 4、监听kafka消息队列，推送页面
 */
func LogFileServiceViewFile(chanName string, filePath string) {

	gm := models.NewGoRoutineManager()
	go gm.TailfFiles(chanName, filePath)

}

func LogFileServiceStopTail(chanName string) {

	gm := models.NewGoRoutineManager()
	err := gm.StopLoopGoroutine(chanName)
	if err != nil {
		common.Logger.Error("gm StopLoopGoroutine failed : %v ", err)
	}

}
