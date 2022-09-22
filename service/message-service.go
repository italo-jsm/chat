package service

import (
	"chat/config"
	"chat/domain"
	"chat/mocking"
	"chat/repository"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)
type MessageService struct{
	Messagerepository repository.MessagerepositoryInterface
}

func NewMessageService(mockRepository bool) (messageService *MessageService){
	newService := MessageService{}
	if(mockRepository){
		newService.Messagerepository = &mocking.MockMessageRepository{}
	}else{
		newService.Messagerepository = &repository.Messagerepository{}
	}
	return &newService
}

func (messageService *MessageService) HandleMessage(message domain.Message){
	message.Consumed = false
	messageService.Messagerepository.SaveMessage(message)
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

func (messageService *MessageService) FindUnreadMessages(userId string) []domain.Message{

	unread := messageService.Messagerepository.FindUnreadMessages(userId)
	var messagesToReturn []domain.Message
	for _, m := range unread{
		m.ConsumeMoment = time.Now()
		m.Consumed = true
		messageService.Messagerepository.UpdateMessage(m)
		messagesToReturn = append(messagesToReturn, m)
	}
	return messagesToReturn
}