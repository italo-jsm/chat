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
	insert, err := database.Prepare("insert into message (id, payload, senderId, receiverId, moment, consumed) values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err.Error())
	}
	result, err2 := insert.Exec(uuid.New().String(), message.Payload, message.SenderId, message.ReceiverId, time.Now().String(), message.Consumed)
	if err2 != nil {
		panic(err2.Error())
	}
	defer database.Close()
	return result
}

func (messageRepository *Messagerepository) FindUnreadMessages(userId string) ([]domain.Message){
	var messages []domain.Message
	database := db.ConnectDatabase()
	rows, err := database.Query("select payload, senderId from message where receiverId = $1 and consumed = false", userId)
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()
	for rows.Next(){
		var m domain.Message
        if err := rows.Scan(&m.Payload, &m.SenderId); err != nil {
            return nil
        }
        messages = append(messages, m)
	}
	defer database.Close()
	return messages
}

func (messageRepository *Messagerepository) UpdateMessage(message domain.Message){
	database := db.ConnectDatabase()
	update, err := database.Prepare("update message set consume_moment = $1, consumed = $2")
	if err != nil {
		panic(err.Error())
	}
	_, err2 := update.Exec(message.ConsumeMoment.String(), message.Consumed)
	if err2 != nil {
		panic(err2.Error())
	}
	defer database.Close()
}
