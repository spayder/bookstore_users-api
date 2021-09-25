package user_services

import (
	"github.com/spayder/bookstore_users-api/domain/users"
	"github.com/spayder/bookstore_users-api/utils/crypto"
	"github.com/spayder/bookstore_users-api/utils/dates"
	"github.com/spayder/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

var (
	UsersService userServiceInterface = &userService{}
)

type userService struct {}


func (u *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.CreatedAt = dates.GetNowString()
	user.Status = StatusActive
	user.Password = crypto.ToMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := u.GetUser(user.Id)
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

func (u *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func (u *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	userDAO := &users.User{}
	return userDAO.FindByStatus(status)
}
