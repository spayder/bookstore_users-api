package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/spayder/bookstore_users-api/utils/errors"
	"strings"
)

func ParseError(err error) *errors.RestErr  {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NotFoundError("No record matching given id")
		}
		return errors.InternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("Invalid data")
	}
	return errors.InternalServerError("Error processing request")
}
