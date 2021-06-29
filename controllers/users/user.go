package users

import (
	"github.com/gin-gonic/gin"
	"github.com/spayder/bookstore_users-api/domain/users"
	"github.com/spayder/bookstore_users-api/services"
	"github.com/spayder/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateHandler(c *gin.Context)  {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	result, resultErr := services.CreateUser(user)

	if resultErr != nil {
		c.JSON(resultErr.Code, resultErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetHandler(c *gin.Context)  {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		userErr := errors.BadRequestError("invalid user id")
		c.JSON(userErr.Code, userErr)
		return
	}

	result, resultErr := services.GetUser(userId)

	if resultErr != nil {
		c.JSON(resultErr.Code, resultErr)
		return
	}

	c.JSON(http.StatusOK, result)
}