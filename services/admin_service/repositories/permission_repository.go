package repositories

import (
	"database/sql"
	"mindmentor/services/admin_service/models"
)

// PermissionRepository представляет интерфейс репозитория для управления разрешениями
type PermissionRepository interface {
	AddPermission(permission models.Permission) error
	UpdatePermission(permission models.Permission) error
	DeletePermission(permissionID int) error
	GetPermissions() ([]models.Permission, error)
}

// SQLPermissionRepository представляет репозиторий для работы с разрешениями в базе данных SQL
type SQLPermissionRepository struct {
	DB *sql.DB
}

// AddPermission добавляет новое разрешение в базу данных SQL
func (r *SQLPermissionRepository) AddPermission(permission models.Permission) error {
	// Подготовка SQL-запроса для вставки разрешения
	query := "INSERT INTO permissions (name, description) VALUES (?, ?)"
	_, err := r.DB.Exec(query, permission.Name, permission.Description)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePermission обновляет разрешение в базе данных
// UpdatePermission обновляет разрешение в базе данных SQL
func (r *SQLPermissionRepository) UpdatePermission(permission models.Permission) error {
	// Подготовка SQL-запроса для обновления разрешения
	query := "UPDATE permissions SET name = ?, description = ? WHERE id = ?"
	_, err := r.DB.Exec(query, permission.Name, permission.Description, permission.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeletePermission удаляет разрешение из базы данных SQL
func (r *SQLPermissionRepository) DeletePermission(permissionID int) error {
	// Подготовка SQL-запроса для удаления разрешения
	query := "DELETE FROM permissions WHERE id = ?"
	_, err := r.DB.Exec(query, permissionID)
	if err != nil {
		return err
	}
	return nil
}

// GetPermissions возвращает список всех разрешений из базы данных SQL
func (r *SQLPermissionRepository) GetPermissions() ([]models.Permission, error) {
	// Подготовка SQL-запроса для выборки всех разрешений
	query := "SELECT id, name, description FROM permissions"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []models.Permission
	for rows.Next() {
		var permission models.Permission
		// Сканирование значений из строк результата запроса в структуру Permission
		err := rows.Scan(&permission.ID, &permission.Name, &permission.Description)
		if err != nil {
			return nil, err
		}
		// Добавление разрешения в список
		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
