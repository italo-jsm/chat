package routes

import (
	"chat/controller"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router{
	messageController := controller.MessageController{}
	userController := controller.UserController{}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/messages", func (w http.ResponseWriter, r *http.Request){
		messageController.SaveMessage(w, r)
	}).Methods("POST")
	router.HandleFunc("/users", func (w http.ResponseWriter, r *http.Request){
		userController.SaveUser(w, r)
	}).Methods("POST")
	router.HandleFunc("/users/{userId}", userController.GetOneUser).Methods("GET")
	router.HandleFunc("/messages/unread/{userId}", messageController.GetUnreadMessages).Methods("GET")
	return router
}