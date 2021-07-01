package users

import (
	"fmt"
	"github.com/spayder/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/spayder/bookstore_users-api/utils/dates"
	"github.com/spayder/bookstore_users-api/utils/errors"
)

var(
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if err := bookstore_users.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]

	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *errors.RestErr  {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("user with an email %s already registered", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("user with an id %d already exists", user.Id))
	}

	user.CreatedAt = dates.GetNowString()
	usersDB[user.Id] = user
	return nil
}
