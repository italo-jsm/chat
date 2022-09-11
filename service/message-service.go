package service

import (
	//"chat/config"
	"chat/config"
	"chat/domain"
	"chat/repository"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleMessage(message domain.Message){
	rep := repository.Messagerepository{}
	message.Consumed = false
	rep.SaveMessage(message)
	if config.GetInstance().NotifyEnabled{
		notifyMessageReceived(message)
	}
}

func notifyMessageReceived(message domain.Message){
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
	p.Flush(1)
}

func FindUnreadMessages(userId string) []domain.Message{
	rep :=repository.Messagerepository{}
	unread := rep.FindUnreadMessages(userId)
	var messagesToReturn []domain.Message
	for _, m := range unread{
		m.ConsumeMoment = time.Now()
		m.Consumed = true
		rep.UpdateMessage(m)
		messagesToReturn = append(messagesToReturn, m)
	}
	return messagesToReturn
}