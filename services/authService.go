package services

import (
	"errors"
	"mini-jira-backend/models"
	"mini-jira-backend/repositories"
	"mini-jira-backend/utils"
)

// login user
func Login(email string, password string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if user == nil {
		return "", errors.New("invalid email or password -- nil user")
	}
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// compare hashed password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(user.ID, user.Email, user.Role)
}

// register user
func Register(user models.User) (models.User, error) {
	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Role = "member"
	user.Password = hashedPassword

	return repositories.RegisterUser(&user)
}
