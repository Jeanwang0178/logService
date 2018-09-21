package models

import (
	"fmt"
	"github.com/hpcloud/tail"
	"logService/src/common"
	"strconv"
	"strings"
	"time"
)

type GoRoutineManager struct {
	grchannelMap *GoroutineChannelMap
}

var (
	tailCount      int32
	RoutineManager *GoRoutineManager
)

func init() {
	gm := &GoroutineChannelMap{}
	RoutineManager = &GoRoutineManager{grchannelMap: gm}
}

func NewGoRoutineManager() *GoRoutineManager {
	if RoutineManager == nil {
		gm := &GoroutineChannelMap{}
		RoutineManager = &GoRoutineManager{grchannelMap: gm}
	}
	return RoutineManager

}

func (gm GoRoutineManager) StopLoopGoroutine(chanName string) error {
	gm.stopOther()
	cacheKey := *new(string)
	common.GetCache(chanName, &cacheKey)
	if cacheKey == "" {
		common.Logger.Error("get cache failed %s ", chanName)
		return fmt.Errorf("get cache failed :" + chanName)
	}
	stopChannel, ok := gm.grchannelMap.grchannels[cacheKey]
	if !ok {
		return fmt.Errorf("not found goroutine name :" + chanName)
	}
	//stopChannel.tails.Done()
	line := tail.Line{"tailf file done ", time.Now(), nil}
	stopChannel.tails.Lines <- &line
	stopChannel.msg <- common.STOP + strconv.Itoa(int(stopChannel.gid))

	return nil
}

func (gm GoRoutineManager) stopOther() error {
	chanName := "4ec3bd45493b4f378630ae630bd579f6"
	cacheKey := *new(string)
	common.GetCache(chanName, &cacheKey)
	if cacheKey == "" {
		common.Logger.Error("get cache failed %s ", chanName)
		return fmt.Errorf("get cache failed :" + chanName)
	}
	stopChannel, ok := gm.grchannelMap.grchannels[cacheKey]
	if !ok {
		return fmt.Errorf("not found goroutine name :" + chanName)
	}
	//stopChannel.tails.Done()
	line := tail.Line{"tailf file done ", time.Now(), nil}
	stopChannel.tails.Lines <- &line
	stopChannel.msg <- common.STOP + strconv.Itoa(int(stopChannel.gid))
	return nil
}
func (gm *GoRoutineManager) NewLoopGoroutine(name string, tails *tail.Tail) {

	go func(this *GoRoutineManager, name string, tails tail.Tail) {
		//register channel
		chanName, err := gm.grchannelMap.register(name, tails)
		if err != nil {
			common.Logger.Error("grchannelMap.register failed %v ", err)
			return
		}
		for {
			select {
			case info := <-this.grchannelMap.grchannels[chanName].msg:
				taskInfo := strings.Split(info, ":")
				signal, gid := taskInfo[0], taskInfo[1]
				if gid == strconv.Itoa(int(this.grchannelMap.grchannels[chanName].gid)) {
					if signal == "_STOP" {

						common.Logger.Info(chanName + "：gid[" + gid + "] quit")
						this.grchannelMap.unregister(chanName)
						common.DeleteCache(name)
						return
					} else {
						common.Logger.Info("unknow signal")
					}
				}
			default:
				//common.Logger.Info("no signal")
			}

			//发送KAFKA消息队列
			msg, ok := <-tails.Lines
			if !ok {
				common.Logger.Info("tail file close reopen, filename:%s\n" + tails.Filename)
				time.Sleep(100 * time.Millisecond)
				return
			}
			err = SendToKafka(msg.Text, common.TopicLog)
			if err != nil {
				common.Logger.Error("taild file error : %v ", err)
			}

		}
	}(gm, name, *tails)
}

func (gm *GoRoutineManager) TailfFiles(chanName string, filePath string) (err error) {

	tails, err := tail.TailFile(filePath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		common.Logger.Error("taild file error : %v ", err)
		return err
	}
	gm.NewLoopGoroutine(chanName, tails)

	return nil
}
