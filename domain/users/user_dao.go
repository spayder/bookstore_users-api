package users

import (
	"fmt"
	"github.com/spayder/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/spayder/bookstore_users-api/utils/dates"
	"github.com/spayder/bookstore_users-api/utils/errors"
	"strings"
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
	stmt, err := bookstore_users.Client.Prepare(
		"INSERT INTO users(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// it's called just before any return of the function
	defer stmt.Close()

	user.CreatedAt = dates.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "email") {
			return errors.BadRequestError(
				fmt.Sprintf("email %s already exists", user.Email),
			)
		}
		return errors.InternalServerError(fmt.Sprintf("Error saving user to database: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("Error saving user to database: %s", err.Error()))
	}

	user.Id = userId
	return nil
}
