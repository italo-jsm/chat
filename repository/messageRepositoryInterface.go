package repository

import (
	"chat/domain"
	"database/sql"
)

type MessagerepositoryInterface interface {
	SaveMessage(message domain.Message) (sql.Result)
	FindUnreadMessages(userId string) ([]domain.Message)
	UpdateMessage(message domain.Message)
}