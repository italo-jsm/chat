package repository

import (
	"time"
	"database/sql"
	"chat/domain"
	"chat/db"
	"github.com/google/uuid"
)

type Messagerepository struct{}

func (messageRepository *Messagerepository) SaveMessage(message domain.Message) (sql.Result){
	database := db.ConnectDatabase()
	insert, err := database.Prepare("insert into message (id, payload, senderId, receiverId, moment) values ($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}
	result, err2 := insert.Exec(uuid.New().String(), message.Payload, message.SenderId, message.ReceiverId, time.Now().String())
	if err2 != nil {
		panic(err2.Error())
	}
	defer database.Close()
	return result
}
