package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spayder/bookstore_users-api/utils/config"
)

var router = gin.Default()

func Handle() {
	MapUrls()
	port := ":" + config.Env("APP_PORT")
	router.Run(port)
}
