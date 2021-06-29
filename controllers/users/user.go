package users

import (
	"github.com/gin-gonic/gin"
	"github.com/spayder/bookstore_users-api/domain/users"
	"github.com/spayder/bookstore_users-api/services"
	"github.com/spayder/bookstore_users-api/utils/errors"
	"net/http"
)

func Create(c *gin.Context)  {
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

func Get(c *gin.Context)  {
	c.String(http.StatusOK, "from users get")
}