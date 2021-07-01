package bookstore_users

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spayder/bookstore_users-api/utils/config"
	"log"
)

const (
	db_username = "DB_USERNAME"
	db_password = "DB_PASSWORD"
	db_host = "DB_HOST"
	db_port = "DB_PORT"
	db_schema = "DB_SCHEMA"
)

var (
	Client *sql.DB
)

func init()  {

	username := config.Env(db_username)
	password := config.Env(db_password)
	host 	 := config.Env(db_host)
	port 	 := config.Env(db_port)
	schema   := config.Env(db_schema)

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