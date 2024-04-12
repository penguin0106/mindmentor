package handlers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"mindmentor/social_service/models"
	"net/http"
)

var db *pgx.Conn // Переменная для подключения к базе данных

func init() {
	// Инициализация соединения с базой данных при запуске приложения
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/database")
	if err != nil {
		panic(err)
	}
}

func CreateChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var chat models.Chat
	_ = json.NewDecoder(r.Body).Decode(&chat)

	// Вставка нового чата в базу данных
	_, err := db.Exec(context.Background(), "INSERT INTO chats (topic) VALUES ($1)", chat.Topic)
	if err != nil {
		http.Error(w, "Failed to create chat", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(chat)
}

func JoinChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var participant models.ChatParticipant
	_ = json.NewDecoder(r.Body).Decode(&participant)

	// Вставка нового участника в чат в базу данных
	_, err := db.Exec(context.Background(), "INSERT INTO chat_participants (chat_id, user_id) VALUES ($1, $2)", participant.ChatID, participant.UserID)
	if err != nil {
		http.Error(w, "Failed to join chat", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(participant)
}
