package handlers

import (
	"encoding/json"
	"log"
	"mindmentor/services/meditation_service/repositories"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
)

// CommentHandler handles HTTP requests related to comments
type CommentHandler struct {
	CommentRepo *repositories.CommentRepository
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

	err = h.CommentRepo.AddComment(&comment)
	if err != nil {
		log.Println("Error adding comment:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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

	comments, err := h.CommentRepo.GetCommentsByCourseID(courseID)
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
