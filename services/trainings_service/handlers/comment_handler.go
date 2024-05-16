package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"trainings_service/models"
	"trainings_service/services"
)

// CommentHandler представляет обработчик HTTP-запросов для работы с комментариями к книгам
type CommentHandler struct {
	CommentServ *services.CommentService
}

// NewCommentHandler создает новый экземпляр обработчика комментариев
func NewCommentHandler(commentServ *services.CommentService) *CommentHandler {
	return &CommentHandler{CommentServ: commentServ}
}

// AddBookCommentHandler обрабатывает запрос на добавление нового комментария к книге
func (h *CommentHandler) AddBookCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("Error decoding comment JSON", err)
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Получение идентификатора пользователя из запроса
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("Error parsing user ID", err)
		http.Error(w, "Некорректный идентификатор пользователя", http.StatusBadRequest)
		return
	}

	// Получение идентификатора книги из запроса
	bookIDStr := r.URL.Query().Get("book_id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		log.Println("Error parsing book ID", err)
		http.Error(w, "Некорректный идентификатор книги", http.StatusBadRequest)
		return
	}

	// Добавление комментария к книге
	err = h.CommentServ.AddBookComment(userID, bookID, comment.Text)
	if err != nil {
		log.Println("Error adding comment:", err)
		http.Error(w, "Ошибка при добавлении комментария", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetBookCommentsHandler обрабатывает запрос на получение комментариев для указанной книги
func (h *CommentHandler) GetBookCommentsHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора книги из запроса
	bookIDStr := r.URL.Query().Get("book_id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		log.Println("Error parsing book ID", err)
		http.Error(w, "Некорректный идентификатор книги", http.StatusBadRequest)
		return
	}

	// Получение комментариев для указанной книги
	comments, err := h.CommentServ.GetBookComments(bookID)
	if err != nil {
		log.Println("Error getting comments for book:", err)
		http.Error(w, "Ошибка при получении комментариев", http.StatusInternalServerError)
		return
	}

	// Отправка комментариев в виде JSON-ответа
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// AddBookRatingHandler обрабатывает запрос на добавление оценки для указанной книги
func (h *CommentHandler) AddBookRatingHandler(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		log.Println("Error decoding rating JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Добавление оценки для указанной книги
	err = h.CommentServ.AddBookRating(&rating)
	if err != nil {
		log.Println("Error adding book rating:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetBookRatingHandler обрабатывает запрос на получение рейтинга книги по ее идентификатору
func (h *CommentHandler) GetBookRatingHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора книги из запроса
	bookIDStr := r.URL.Query().Get("book_id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		log.Println("Error parsing book ID", err)
		http.Error(w, "Некорректный идентификатор книги", http.StatusBadRequest)
		return
	}

	// Получение рейтинга книги по ее идентификатору
	rating, err := h.CommentServ.GetBookRating(bookID)
	if err != nil {
		log.Println("Error getting book rating:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправка рейтинга книги в виде JSON-ответа
	response := map[string]float64{"rating": rating}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}
