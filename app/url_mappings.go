package app

import (
	"github.com/spayder/bookstore_users-api/controllers/ping"
	"github.com/spayder/bookstore_users-api/controllers/users"
)

func MapUrls()  {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetHandler)
	router.POST("/users", users.CreateHandler)
}
