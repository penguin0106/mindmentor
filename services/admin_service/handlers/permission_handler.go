package handlers

import (
	"encoding/json"
	"fmt"
	"mindmentor/services/admin_service/models"
	"mindmentor/services/admin_service/repositories"
	"net/http"
	"strconv"
)

// PermissionHandler представляет обработчик запросов для работы с разрешениями
type PermissionHandler struct {
	Repository repositories.PermissionRepository
}

// NewPermissionHandler создает новый экземпляр PermissionHandler
func NewPermissionHandler(repo repositories.PermissionRepository) *PermissionHandler {
	return &PermissionHandler{Repository: repo}
}

// AddPermissionHandler обрабатывает запрос на добавление разрешения
func (h *PermissionHandler) AddPermissionHandler(w http.ResponseWriter, r *http.Request) {
	var permission models.Permission
	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		http.Error(w, "Ошибка при декодировании данных запроса", http.StatusBadRequest)
		return
	}

	err = h.Repository.AddPermission(permission)
	if err != nil {
		http.Error(w, "Ошибка при добавлении разрешения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Разрешение успешно добавлено")
}

// UpdatePermissionHandler обрабатывает запрос на обновление разрешения
func (h *PermissionHandler) UpdatePermissionHandler(w http.ResponseWriter, r *http.Request) {
	var permission models.Permission
	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		http.Error(w, "Ошибка при декодировании данных запроса", http.StatusBadRequest)
		return
	}

	err = h.Repository.UpdatePermission(permission)
	if err != nil {
		http.Error(w, "Ошибка при обновлении разрешения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Разрешение успешно обновлено")
}

// DeletePermissionHandler обрабатывает запрос на удаление разрешения
func (h *PermissionHandler) DeletePermissionHandler(w http.ResponseWriter, r *http.Request) {
	permissionIDStr := r.URL.Query().Get("id")
	permissionID, err := strconv.Atoi(permissionIDStr)
	if err != nil {
		http.Error(w, "ID разрешения должен быть числом", http.StatusBadRequest)
		return
	}

	err = h.Repository.DeletePermission(permissionID)
	if err != nil {
		http.Error(w, "Ошибка при удалении разрешения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Разрешение успешно удалено")
}

// GetPermissionsHandler обрабатывает запрос на получение списка разрешений
func (h *PermissionHandler) GetPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.Repository.GetPermissions()
	if err != nil {
		http.Error(w, "Ошибка при получении списка разрешений", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}
