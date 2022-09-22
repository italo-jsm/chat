package controller

import (
	"chat/domain"
	"chat/service"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type MessageController struct{}

func (messageController *MessageController) SaveMessage(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	messageService := service.NewMessageService(false)
	if err != nil {
		panic(err.Error)
	}

	message := domain.Message{}
	json.Unmarshal(reqBody, &message)
	messageService.HandleMessage(message)
	w.Write([]byte(http.StatusText(200)))
}

func (MessageController *MessageController) GetUnreadMessages(w http.ResponseWriter, r *http.Request){
	messageService := service.NewMessageService(false)
	json.NewEncoder(w).Encode(messageService.FindUnreadMessages(mux.Vars(r)["userId"]))
}