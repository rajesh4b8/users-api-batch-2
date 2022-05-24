package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rajesh4b8/users-api-batch-2/src/logger"
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
		logger.Error("DB connection failed", err)
		panic(err)
	}

	logger.Info("Database connection successful")
}
