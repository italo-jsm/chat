package controller

import (
	"chat/domain"
	"chat/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct{}

func (userController *UserController) SaveUser(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error)
	}

	user := domain.User{}
	json.Unmarshal(reqBody, &user)
	rep := repository.UserRepository{}
	rep.SaveUser(user)
	w.Write([]byte(http.StatusText(200)))
}

func (userController *UserController) GetOneUser(w http.ResponseWriter, r *http.Request){
	userId := mux.Vars(r)["userId"]
	rep := repository.UserRepository{}
	user := rep.FindOneUser(userId)
	if (user == nil){
		w.WriteHeader(http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(user)
}