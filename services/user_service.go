package services

import (
	"github.com/spayder/bookstore_users-api/domain/users"
	"github.com/spayder/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User)  (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := users.User{Id: userId}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}