package dao

import (
	"errors"

	"github.com/pgillich/meals-demo/internal/models"
)

var (
	ErrUnknownUser = errors.New("unknown user")
	ErrBadPassword = errors.New("bad password")
)

func (dbHandler *Handler) AuthenticateUser(email string, password string) (*models.User, error) {
	if email != "yoda@star.wars" {
		return nil, ErrUnknownUser
	}
	if password != "master" {
		return nil, ErrBadPassword
	}

	return dbHandler.GetUserByEmail(email)
}

func (dbHandler *Handler) GetUserByEmail(email string) (*models.User, error) {
	if email != "yoda@star.wars" {
		return nil, ErrUnknownUser
	}

	return &models.User{
		Email:    email,
		FullName: "Yoda Master",
	}, nil
}
