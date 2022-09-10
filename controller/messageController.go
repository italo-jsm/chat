package controller

import (
	"chat/domain"
	"chat/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type MessageController struct{}

func (messageController *MessageController) SaveMessage(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error)
	}

	message := domain.Message{}
	json.Unmarshal(reqBody, &message)
	service.HandleMessage(message)
	w.Write([]byte(http.StatusText(200)))
}