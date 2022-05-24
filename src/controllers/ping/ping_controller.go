package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "pong",
	// })

	c.String(http.StatusOK, "Pong")
}
