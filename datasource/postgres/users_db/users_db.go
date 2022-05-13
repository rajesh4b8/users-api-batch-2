package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	username = "postgres"
	host     = "127.0.0.1"
	schema   = "postgres"
	password = "dev"
)

var (
	Client *sql.DB
)

func init() {
	datasource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username,
		password,
		host,
		schema,
	)
	var error error
	Client, error = sql.Open("postgres", datasource)
	if error != nil {
		panic(error)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database connection successful!!")
}
