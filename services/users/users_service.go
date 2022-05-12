package users

import (
	"github.com/rajesh4b8/users-api-batch-2/domain/users"
	"github.com/rajesh4b8/users-api-batch-2/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{
		Id: userId,
	}

	if restErr := user.Get(); restErr != nil {
		return nil, restErr
	}

	return user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if restErr := user.Save(); restErr != nil {
		return nil, restErr
	}

	return &user, nil
}
