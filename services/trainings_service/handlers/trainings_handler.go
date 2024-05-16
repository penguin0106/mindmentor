package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trainings_service/services"
)

// TrainingHandler представляет обработчик HTTP-запросов для работы с тренировками
type TrainingHandler struct {
	TrainingServ *services.TrainingService
}

func NewTrainingHandler(trainingServ *services.TrainingService) *TrainingHandler {
	return &TrainingHandler{TrainingServ: trainingServ}
}

func (h *TrainingHandler) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Content     []byte `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	err = h.TrainingServ.AddBook(reqBody.Title, reqBody.Description, reqBody.Content)
	if err != nil {
		http.Error(w, "Ошибка при добавлении книги: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TrainingHandler) GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "Не указан идентификатор книги", http.StatusBadRequest)
		return
	}

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, "Некорректный формат идентификатора книги", http.StatusBadRequest)
		return
	}

	book, err := h.TrainingServ.GetBookByID(bookIDInt)
	if err != nil {
		http.Error(w, "Ошибка при получении книги: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}

	// Отправляем книгу в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *TrainingHandler) GetAllBooksHandler(w http.ResponseWriter, _ *http.Request) {
	books, err := h.TrainingServ.GetAllBooks()
	if err != nil {
		http.Error(w, "Ошибка при получении списка книг: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем список книг в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *TrainingHandler) GetBookByNameHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Не указано название книги", http.StatusBadRequest)
		return
	}

	book, err := h.TrainingServ.GetBookByName(title)
	if err != nil {
		http.Error(w, "Ошибка при получении книги: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}

	// Отправляем книгу в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
