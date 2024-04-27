package handlers

import (
	"encoding/json"
	"log"
	"mindmentor/services/meditation_service/repositories"
	"net/http"
)

// CourseHandler handles HTTP requests related to courses
type CourseHandler struct {
	CourseRepo *repositories.CourseRepository
}

// GetAllCoursesHandler returns all meditation courses
func (h *CourseHandler) GetAllCoursesHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := h.CourseRepo.GetAllCourses()
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
