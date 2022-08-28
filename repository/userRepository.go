package repository

import (
	"chat/db"
	"chat/domain"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository struct{}

func (userRepository *UserRepository) SaveUser(user domain.User) (sql.Result){
	database := db.ConnectDatabase()
	insert, err := database.Prepare("insert into chat_user (id, username, email, public_key) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	result, err2 := insert.Exec(uuid.New().String(), user.Username, user.Email, user.PublicKey)
	if err2 != nil {
		panic(err2.Error())
	}
	defer database.Close()
	return result
}

func (userRepository *UserRepository) FindOneUser(id string) *domain.User {
	fmt.Println("Buscando user com id: " + id)
	database := db.ConnectDatabase()
	sqlStatement := "select * from chat_user where id=$1"
	row := database.QueryRow(sqlStatement, id)
	var user domain.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PublicKey)
	if err != nil {
		return nil
	}
	defer database.Close()
	return &user
}