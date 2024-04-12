package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mindmentor/social_service/handlers"
)

func main() {
	router := mux.NewRouter()

	// Установка обработчиков для эндпоинтов создания чата и присоединения к чату
	router.HandleFunc("/chat", handlers.CreateChat).Methods("POST")
	router.HandleFunc("/join", handlers.JoinChat).Methods("POST")

	log.Fatal(http.ListenAndServe(":8084", router))
}
