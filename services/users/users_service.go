package users

import (
	"errors"

	"github.com/rajesh4b8/users-api-batch-2/domain/users"
)

func GetUser(userId int64) (*users.User, error) {
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

	return nil, errors.New("user not found")
}
