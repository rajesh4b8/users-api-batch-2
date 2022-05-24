package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rajesh4b8/users-api-batch-2/src/domain/users"
	"github.com/rajesh4b8/users-api-batch-2/src/services"
	"github.com/rajesh4b8/users-api-batch-2/src/utils/errors"
)

var (
	userService = services.UsersService
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

	user, restErr := userService.GetUser(userId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	restErr := errors.NewBadRequestError("Input not readable")
	// 	c.JSON(restErr.Status, restErr)
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err)
	// 	restErr := errors.NewBadRequestError("Invalid json")
	// 	c.JSON(restErr.Status, restErr)
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.NewBadRequestError("Invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}

	if restErr := user.Validate(); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := userService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func FetchAllUsers(c *gin.Context) {

	// should return a array of Users in JSON format

}
