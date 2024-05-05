package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"meditation_service/models"
	"meditation_service/services"
	"net/http"
	"strconv"
	"time"
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
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("Error decoding comment JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//Проверка наличия необходимых полей в комментарии
	if comment.UserID == 0 || comment.Text == "" {
		http.Error(w, "Недостаточно данных для добавления комментария", http.StatusBadRequest)
	}

	//Установка временной метки комментария
	comment.Timestamp = time.Now().Unix()

	err = h.CommentService.AddCourseComment(&comment)
	if err != nil {
		log.Println("Ошибка добавления комментария:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Комментарий успешно добавлен")
}

// GetCommentsByCourseIDHandler returns all comments for a given course ID
func (h *CommentHandler) GetCommentsByCourseIDHandler(w http.ResponseWriter, r *http.Request) {
	courseIDStr := r.URL.Query().Get("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		log.Println("Error parsing course ID:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	comments, err := h.CommentService.GetCourseComments(courseID)
	if err != nil {
		log.Println("Error getting comments by course ID:", err)
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
