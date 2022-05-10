package app

import (
	"github.com/rajesh4b8/users-api-batch-2/controllers/ping"
	"github.com/rajesh4b8/users-api-batch-2/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
