package service

import (
	"errors"

	"github.com/sumitst05/patiently/internal/models"
	"github.com/sumitst05/patiently/internal/repository"
	"github.com/sumitst05/patiently/utils"
)

func RegisterUser(name, email, password, role string) (*models.User, error) {
	if role != models.RoleReceptionist && role != models.RoleDoctor {
		return nil, errors.New("Invalid role")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		Role:         role,
	}

	return repository.CreateUser(user)
}

func AuthenticateUser(email, password string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("Invalid credentials")
	}

	return user, nil
}
