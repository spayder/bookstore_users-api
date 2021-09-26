package users

import (
	"github.com/spayder/bookstore_users-api/datasources/mysql/bookstore_users"
	"github.com/spayder/bookstore_users-api/logger"
	"github.com/spayder/bookstore_users-api/utils/dates"
	"github.com/spayder/bookstore_users-api/utils/errors"
)

func (user *User) Get() *errors.RestErr {
	query := "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE id = ?;"
	stmt, err := bookstore_users.Client.Prepare(query)

	if err != nil {
		logger.Error("error trying to prepare Get User statement", err)
		return errors.InternalServerError("a database error occurred")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
		logger.Error("error trying to Get a User by ID statement", err)
		return errors.InternalServerError("a database error occurred")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := bookstore_users.Client.Prepare(
		"INSERT INTO users(first_name, last_name, email, created_at, status, password) VALUES (?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		logger.Error("error trying to prepare Save User statement", err)
		return errors.InternalServerError("a database error occurred")
	}
	// it's called just before any return of the function
	defer stmt.Close()

	user.CreatedAt = dates.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.Status, user.Password)
	if err != nil {
		logger.Error("error trying to execute Save User statement", err)
		return errors.InternalServerError("a database error occurred")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error trying to Get Last Inserted Id after creating a new user", err)
		return errors.InternalServerError("a database error occurred")
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	query := "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"

	stmt, err := bookstore_users.Client.Prepare(query)
	if err != nil {
		logger.Error("error trying to prepare Update User statement", err)
		return errors.InternalServerError("a database error occurred")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error trying to execute Update User", err)
		return errors.InternalServerError("a database error occurred")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	query := "DELETE FROM users WHERE id = ?"

	stmt, err := bookstore_users.Client.Prepare(query)
	if err != nil {
		logger.Error("error trying to prepare Delete User statement", err)
		return errors.InternalServerError("a database error occurred")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error("error trying to execute Delete User statement", err)
		return errors.InternalServerError("a database error occurred")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	query := "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE status = ?"

	stmt, err := bookstore_users.Client.Prepare(query)
	if err != nil {
		logger.Error("error trying to prepare Find Users By Status statement", err)
		return nil, errors.InternalServerError("a database error occurred")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error trying to execute Find Users By Status statement", err)
		return nil, errors.InternalServerError("a database error occurred")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
			logger.Error("error trying to scan row for Find Users By Status statement", err)
			return nil, errors.InternalServerError("a database error occurred")

		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFoundError("no users matching given criteria")
	}
	return results, nil
}
