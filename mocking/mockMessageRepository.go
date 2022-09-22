package mocking

import (
	"chat/domain"
	"database/sql"
)

type MockMessageRepository struct{}

func (messageRepository *MockMessageRepository) SaveMessage(message domain.Message) (sql.Result){
	return nil
}

func (messageRepository *MockMessageRepository) FindUnreadMessages(userId string) ([]domain.Message){
	messages := make([]domain.Message, 1)
	message := domain.Message{
		Payload: "payload",
		SenderId: "senderId",
	}
	messages[0] = message
	return messages
}

func (messageRepository *MockMessageRepository) UpdateMessage(message domain.Message){
}