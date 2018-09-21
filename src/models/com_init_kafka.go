package models

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"logService/src/common"
	"strings"
	"sync"
)

var (
	produce       sarama.SyncProducer
	consumer      sarama.Consumer
	partitionList []int32
	wg            sync.WaitGroup
)

func InitKafka() (err error) {

	err = initProduce()
	if err != nil {
		return err
	}

	return
}

// 初始化KAFKA生产者
func initProduce() (err error) {
	//初始化KAFKA配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//创建生产者
	kafkaServer := beego.AppConfig.String("kafka.producer.servers")
	servers := strings.Split(strings.TrimSpace(kafkaServer), ",")
	produce, err = sarama.NewSyncProducer(servers, config)

	if err != nil {
		common.Logger.Error("sarama.NewSyncProducer failed ", err)
		return
	}
	return err
}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := produce.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v pid:%v offset:%v data:%v topic:%v", err, pid, offset, data, topic)
		return
	}

	return
}
