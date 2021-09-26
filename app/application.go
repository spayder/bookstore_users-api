package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spayder/bookstore_users-api/logger"
	"github.com/spayder/bookstore_users-api/utils/config"
)

var router = gin.Default()

func Handle() {
	MapUrls()

	logger.Info("about to start the application ...")
	port := ":" + config.Env("APP_PORT")
	logger.Info(fmt.Sprintf("application is listening on port %s", port))
	router.Run(port)
}
