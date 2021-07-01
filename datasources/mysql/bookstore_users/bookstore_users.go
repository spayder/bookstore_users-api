package bookstore_users

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spayder/bookstore_users-api/utils/config"
	"log"
)

var (
	Client *sql.DB
)

func init()  {
	username := config.Env("DB_USERNAME")
	password := config.Env("DB_PASSWORD")
	host 	 := config.Env("DB_HOST")
	port 	 := config.Env("DB_PORT")
	schema   := config.Env("DB_SCHEMA")

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		schema,
	)

	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully connected")
}