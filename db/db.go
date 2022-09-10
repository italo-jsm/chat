package db

import (
	"chat/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

  func ConnectDatabase() *sql.DB{
	conf := config.GetInstance()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
  }