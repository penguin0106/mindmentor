package handlers

import (
	"encoding/json"
	"meditation_service/models"
	"meditation_service/services"
	"net/http"
)

type FavouriteHandler struct {
	FavService *services.FavoriteService
}

func NewFavoriteHandler(favService *services.FavoriteService) *FavouriteHandler {
	return &FavouriteHandler{
		FavService: favService,
	}
}

// AddToFavouriteHandler обрабатывает запрос на добавление элемента в избранное
func (h *FavouriteHandler) AddToFavouriteHandler(w http.ResponseWriter, r *http.Request) {
	var favorite *models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = h.FavService.AddToFavourite(favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RemoveFromFavoriteHandler обрабатывает запрос на удаление элемента из избранного
func (h *FavouriteHandler) RemoveFromFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	var favorite *models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = h.FavService.RemoveFromFavorite(favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
