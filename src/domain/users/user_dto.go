package users

import (
	"strings"

	"github.com/rajesh4b8/users-api-batch-2/src/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" || !strings.Contains(user.Email, "@") {
		return errors.NewBadRequestError("Invalid Email id")
	}

	return nil
}
