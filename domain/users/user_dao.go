package users

import (
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/rajesh4b8/users-api-batch-2/datasource/postgres/users_db"
	"github.com/rajesh4b8/users-api-batch-2/logger"
	"github.com/rajesh4b8/users-api-batch-2/utils/date_utils"
	"github.com/rajesh4b8/users-api-batch-2/utils/errors"
)

const (
	noRowsError = "no rows in result set"
	// emailAlreadyExistsError = "duplicate key value violates unique constraint \"unique_email\""
	queryFetchUser        = "select user_id, first_name, last_name, email, date_created from users where user_id = $1"
	queryInsertUser       = "insert into users(first_name, last_name, email) values ($1, $2, $3) RETURNING user_id, date_created"
	queryFetchUserByEmail = "select user_id, first_name, last_name, email, date_created from users where email = $1"

	errorCodeUniqueConstraint = "23505"
)

var (
// usersDB = make(map[int64]*User)
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
		logger.Error("Error while prepare stmt the user", err)
		return errors.NewInternalServerError("Database error")
	}
	// defer the execution to close stmt until before return
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {

		if strings.Contains(err.Error(), noRowsError) {
			return errors.NewNotFoundError(fmt.Sprintf("User not found for id %d", user.Id))
		}

		logger.Error(fmt.Sprintf("Error while fetching user %d", user.Id), err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFetchUserByEmail)
	if err != nil {
		logger.Error("Error while prepare stmt the user", err)
		return errors.NewInternalServerError("Database error")
	}
	// defer the execution to close stmt until before return
	defer stmt.Close()

	if err := stmt.QueryRow(user.Email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), noRowsError) {
			return nil
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error while fetching user %d", user.Id))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	// if usersDB[user.Id] != nil {
	// 	return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	// }

	// user.DateCreated = date_utils.GetNowString()
	// usersDB[user.Id] = user
	if user.Id != 0 {
		return errors.NewBadRequestError("User not supposed to porovide user_id")
	}

	// One way to validate is by fetching the user by email
	// if err := user.GetByEmail(); err != nil {
	// 	return errors.NewInternalServerError("Error while searching user by email")
	// }

	// if user.Id != 0 {
	// 	return errors.NewBadRequestError("Email already exists")
	// }

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	defer stmt.Close()
	if err != nil {
		return errors.NewInternalServerError("Error while preparing DB stmt")
	}
	var dateCreated time.Time
	if err := stmt.QueryRow(user.FirstName, user.LastName, user.Email).Scan(&user.Id, &dateCreated); err != nil {
		sqlErr, ok := err.(*pq.Error)
		if !ok {
			return errors.NewInternalServerError("Error while reading the sql error")
		}

		if sqlErr.Code == errorCodeUniqueConstraint {
			return errors.NewNotFoundError(fmt.Sprintf("User already exists with email %v", user.Email))
		}
		return errors.NewInternalServerError("Error while insert ::" + err.Error())
	}
	user.DateCreated = dateCreated.Format(date_utils.ApiDateFormat)

	return nil
}

func GetAllUsers() ([]User, *errors.RestErr) {
	// TODO for Assignment
	// Get all the users from map
	// if map is empty return 404 saying List is empty

	return nil, errors.NewNotFoundError("No users found")
}
