package handlers

import (
	"encoding/json"
	"log"
	"meditation_service/models"
	"meditation_service/services"
	"net/http"
)

// CourseHandler handles HTTP requests related to courses
type CourseHandler struct {
	CourseServ *services.CourseService
}

func NewCourseHandler(courServ *services.CourseService) *CourseHandler {
	return &CourseHandler{CourseServ: courServ}
}

// GetAllCoursesHandler returns all meditation courses
func (h *CourseHandler) GetAllCoursesHandler(w http.ResponseWriter, _ *http.Request) {
	courses, err := h.CourseServ.GetAllCourses()
	if err != nil {
		log.Println("Error getting all courses:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(courses)
	if err != nil {
		log.Println("Error marshalling courses to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetCourseByNameHandler обрабатывает запрос на получение курса медитации по его названию
func (h *CourseHandler) GetCourseByNameHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра courseName из запроса
	courseName := r.URL.Query().Get("courseName")

	// Получение курса медитации по его названию с использованием CourseService
	course, err := h.CourseServ.GetCourseByName(courseName)
	if err != nil {
		// Обработка ошибки
		http.Error(w, "Failed to get course by name", http.StatusInternalServerError)
		return
	}
	if course == nil {
		// Обработка ситуации, когда курс не найден
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	// Отправка информации о курсе в формате JSON
	jsonResponse(w, course)
}

// AddCourseHandler обрабатывает запрос на добавление нового курса медитации
func (h *CourseHandler) AddCourseHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение данных курса из тела запроса
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		// Обработка ошибки
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Добавление нового курса медитации с использованием CourseService
	err = h.CourseServ.AddCourse(&course)
	if err != nil {
		// Обработка ошибки
		http.Error(w, "Failed to add course", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
}

// jsonResponse отправляет JSON-ответ клиенту
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// Обработка ошибки
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}
