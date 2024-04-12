package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mindmentor/meditation_service/handlers"
)

func main() {
	router := mux.NewRouter()

	// Установка обработчика для эндпоинта получения информации о практиках медитации
	router.HandleFunc("/meditations", handlers.GetMeditations).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", router))
}
