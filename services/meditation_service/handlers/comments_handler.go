package handlers

import (
	"encoding/json"
	"log"
	"meditation_service/services"
	"net/http"
	"strconv"
)

// CommentHandler handles HTTP requests related to comments
type CommentHandler struct {
	CommentService *services.CommentService
}

func NewCommentHandler(comServ *services.CommentService) *CommentHandler {
	return &CommentHandler{CommentService: comServ}
}

// AddCommentHandler adds a new comment for a course
func (h *CommentHandler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	//Извлекаем данные из запроса
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		http.Error(w, "Некорректный идентификатор пользователя", http.StatusBadRequest)
		return
	}

	itemId, err := strconv.Atoi(r.FormValue("item_id"))
	if err != nil {
		http.Error(w, "Некорректный идентификатор элемента", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Текст комментария не может быть пустым", http.StatusBadRequest)
		return
	}

	err = h.CommentService.AddVideoComment(userID, itemId, text)
	if err != nil {
		http.Error(w, "Ошибка при добавлении комментария", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetCommentsByVideoIDHandler returns all comments for a given course ID
func (h *CommentHandler) GetCommentsByVideoIDHandler(w http.ResponseWriter, r *http.Request) {
	videoIDStr := r.URL.Query().Get("item_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		log.Println("Error parsing video ID:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	comments, err := h.CommentService.GetVideoComments(videoID)
	if err != nil {
		log.Println("Error getting comments by video ID:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(comments)
	if err != nil {
		log.Println("Error marshalling comments to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
