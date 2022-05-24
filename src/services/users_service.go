package services

import (
	"github.com/rajesh4b8/users-api-batch-2/src/domain/users"
	"github.com/rajesh4b8/users-api-batch-2/src/utils/errors"
)

var (
	UsersService usersServiceinterface = &usersService{}
)

type usersService struct {
}

type usersServiceinterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{
		Id: userId,
	}

	if restErr := user.Get(); restErr != nil {
		return nil, restErr
	}

	return user, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if restErr := user.Save(); restErr != nil {
		return nil, restErr
	}

	return &user, nil
}
