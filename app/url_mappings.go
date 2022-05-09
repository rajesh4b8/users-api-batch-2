package app

import (
	"github.com/rajesh4b8/users-api-batch-2/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
