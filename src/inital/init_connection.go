package inital

import (
	"logService/src/common"
	"logService/src/models"
)

func Init() {

	common.InitCache()

	models.InitKafka()

}
