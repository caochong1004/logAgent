package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var client sarama.SyncProducer
func Init() (err error){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true
	// 构造一个消息

	client, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	return err
}

func SendToMessage(topic, msg string)  {
	message := &sarama.ProducerMessage{}
	message.Topic = topic
	message.Value = sarama.StringEncoder(msg)
	pid, offset, err := client.SendMessage(message)
	if err != nil{
		fmt.Println("send msg failed")
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
