package service

import (
	//"chat/config"
	"chat/domain"
	"chat/repository"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleMessage(message domain.Message){
	rep := repository.Messagerepository{}
	rep.SaveMessage(message)
	notifyMessageReceived(message)
}

func notifyMessageReceived(message domain.Message){
	//conf := config.GetInstance()
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := "messages"

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message.ReceiverId),
	}, nil)
	p.Flush(5000)
}