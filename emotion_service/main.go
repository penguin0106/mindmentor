package main

import (
	"github.com/gorilla/mux"
	"log"
	"mindmentor/emotion_service/handlers"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/emotions", handlers.GetEmotions).Methods("GET")
	router.HandleFunc("/emotions", handlers.AddEmotion).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}
