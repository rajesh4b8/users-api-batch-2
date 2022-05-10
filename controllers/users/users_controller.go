package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rajesh4b8/users-api-batch-2/services/users"
	"github.com/rajesh4b8/users-api-batch-2/utils/errors"
)

func GetUser(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		// response a bad request saying user_id is not a valid number
		restErr := errors.NewBadRequestError("user_id not a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := users.GetUser(userId)
	if getErr != nil {
		restErr := errors.NewNotFoundError(fmt.Sprintf("user not found for id %d", userId))
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

}
