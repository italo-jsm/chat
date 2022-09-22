package tests

import (
	"chat/service"
	"testing"
)

func TestMessageService(t *testing.T) {
	messageService := service.NewMessageService(true)
	receivedMessages := messageService.FindUnreadMessages("id")
	if(receivedMessages[0].Consumed == false){
		t.Fail()
	}
}