package users

import (
	"github.com/gin-gonic/gin"
	"github.com/spayder/bookstore_users-api/domain/users"
	"github.com/spayder/bookstore_users-api/services"
	"github.com/spayder/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.BadRequestError("invalid user id")
	}

	return userId, nil
}

func CreateHandler(c *gin.Context) {
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

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetHandler(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	result, resultErr := services.GetUser(userId)
	if resultErr != nil {
		c.JSON(resultErr.Code, resultErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateHandler(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	user.Id = userId

	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteHandler(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"success": "true"})
}

func SearchHandler(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
