package users

import (
	"time"

	"github.com/rajesh4b8/users-api-batch-2/domain/users"
	"github.com/rajesh4b8/users-api-batch-2/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := users.User{
		Id:          123,
		FirstName:   "Dummy First",
		LastName:    "Dummy Last",
		Email:       "dummy-test@test.com",
		DateCreated: "Something",
	}
	if userId == 123 {
		return &user, nil
	}

	return nil, errors.NewNotFoundError("user not found")
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.DateCreated = time.Now().String()

	return &user, nil
}
