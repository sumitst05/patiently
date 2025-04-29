package repository

import (
	"fmt"

	"github.com/sumitst05/patiently/internal/models"
)

func CreateUser(user *models.User) (*models.User, error) {
	// validate roles received from payload
	if user.Role != models.RoleReceptionist && user.Role != models.RoleDoctor {
		return nil, fmt.Errorf("Invalid role: %s", user.Role)
	}

	err := DB.Create(user).Error
	return user, err
}

func GetUserById(id uint) (*models.User, error) {
	user := models.User{}
	err := DB.First(&user, id).Error

	return &user, err
}

func GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := DB.Where("email = ?", email).First(&user).Error

	return &user, err
}
