package users

import (
	"github.com/spayder/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/spayder/bookstore_users-api/utils/dates"
	"github.com/spayder/bookstore_users-api/utils/errors"
	"github.com/spayder/bookstore_users-api/utils/mysql"
)

func (user *User) Get() *errors.RestErr {
	query := "SELECT id, first_name, last_name, email FROM users WHERE id = ?;"
	stmt, err := bookstore_users.Client.Prepare(query)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return mysql.ParseError(err)
	}

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
		return mysql.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql.ParseError(err)
	}

	user.Id = userId
	return nil
}
