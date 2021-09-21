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
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		current.FirstName = user.FirstName
	}

	if user.LastName != "" {
		current.LastName = user.LastName
	}

	if user.Email != "" {
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}