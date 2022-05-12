package users

import (
	"fmt"

	"github.com/rajesh4b8/users-api-batch-2/utils/date_utils"
	"github.com/rajesh4b8/users-api-batch-2/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	current := usersDB[user.Id]

	if current == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.FirstName = current.FirstName
	user.LastName = current.LastName
	user.Email = current.Email
	user.DateCreated = current.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user

	return nil
}

func GetAllUsers() ([]User, *errors.RestErr) {
	// TODO for Assignment
	// Get all the users from map
	// if map is empty return 404 saying List is empty

	return nil, errors.NewNotFoundError("No users found")
}
