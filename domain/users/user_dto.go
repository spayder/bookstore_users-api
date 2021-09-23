package users

import (
	"github.com/spayder/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.BadRequestError("invalid email address")
	}
	if user.Password == "" {
		return errors.BadRequestError("invalid password")
	}
	return nil
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))

	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}

	return result
}
