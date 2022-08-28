package main

import (
	"chat/routes"
	"log"
	"net/http"
)

func main(){
	router := routes.CreateRoutes()
	log.Fatal(http.ListenAndServe(":3000", router))
}