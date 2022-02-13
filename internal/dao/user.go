package dao

import (
	"fmt"

	"github.com/pgillich/meals-demo/internal/models"
)

func (dbHandler *Handler) AuthenticateUser(email string, password string) (*models.User, error) {
	if email != "yoda@star.wars" {
		return nil, fmt.Errorf("unkown user")
	}
	if password != "master" {
		return nil, fmt.Errorf("bad password")
	}

	return dbHandler.GetUserByEmail(email)
}

func (dbHandler *Handler) GetUserByEmail(email string) (*models.User, error) {
	if email != "yoda@star.wars" {
		return nil, fmt.Errorf("unkown user")
	}

	return &models.User{
		Email:    email,
		FullName: "Yoda Master",
	}, nil
}
