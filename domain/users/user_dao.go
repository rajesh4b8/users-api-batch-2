package users

import (
	"fmt"
	"strings"

	"github.com/rajesh4b8/users-api-batch-2/datasource/postgres/users_db"
	"github.com/rajesh4b8/users-api-batch-2/utils/date_utils"
	"github.com/rajesh4b8/users-api-batch-2/utils/errors"
)

const (
	noRowsError    = "no rows in result set"
	queryFetchUser = "select user_id, first_name, last_name, email, date_created from users where user_id = $1"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		return errors.NewInternalServerError("Database connection failed")
	}

	// current := usersDB[user.Id]

	// if current == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	// }

	// user.FirstName = current.FirstName
	// user.LastName = current.LastName
	// user.Email = current.Email
	// user.DateCreated = current.DateCreated

	stmt, err := users_db.Client.Prepare(queryFetchUser)
	if err != nil {
		return errors.NewInternalServerError("Error while preparing DB stmt")
	}
	// defer the execution to close stmt until before return
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NewNotFoundError(fmt.Sprintf("User not found for id %d", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error while fetching user %d", user.Id))
	}

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
