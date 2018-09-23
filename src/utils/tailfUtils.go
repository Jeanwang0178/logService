package utils

import (
	"logService/src/models"
)

var (
	tailCount int32
)

func TailfFiles(gm *models.GoRoutineManager) {

	/*fileName := "C:\\data\\logs\\sinochem-oms.log"
	tails, err := tail.TailFile(fileName, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		common.Logger.Error("taild file error : %v ", err)
	}

	gm.NewLoopGoroutine("", tails)
	*/
	return
}
