package repositories

import (
	"database/sql"
	"errors"
	"mini-jira-backend/configs"
	"mini-jira-backend/models"
)

// get user by user id
func GetUserByID(id int) (*models.User, error) {
	row := configs.DB.QueryRow(`
		SELECT id, name, email, password, role, created_at
		FROM users WHERE id = $1
	`, id)
	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if no user found
		}
		return nil, err
	}
	return &user, nil
}

func RegisterUser(user *models.User) (models.User, error) {
	err := configs.DB.QueryRow(`
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)

		RETURNING id
	`, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID)
	return *user, err
}

// get user by email
func GetUserByEmail(email string) (*models.User, error) {
	row := configs.DB.QueryRow(`
		SELECT id, name, email, password, role, created_at
		FROM users WHERE email = $1
	`, email)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found") // ✅ FIX
		}
		return nil, err
	}

	return &user, nil
}
