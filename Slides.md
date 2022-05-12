# Notes

## Environment setup

- Define GOROOT -> Should done already when Go is installed
- Define GOPATH -> Default should ~/go or you can define your own
- Clone the repository
- Run the code and see how it works!

## Basic application setup

- Define router in application.go
- Deinne mappings for routes (ex: /ping) in mappings.go
- Start your web app from application.go with router.Run
- Defined our first controller ping_controller.go
- Attached the route /ping to the func Ping from ping_controller.go

## Create User

- Bind the request body to User struct
- Validate the email id
- Pass the user to user service to create user
- Call the Dao to save the user
- Error handling the user already exists

## Get User

- Get the user_id as path param
- Validate the user_id is number
- Pass the user_id to service
- Fetch the user based on user id from Dao
- Error if user not found

## Fetch all users (Assignment)

- A new GET endpoint to return list of Users
- Create new function in service to fetch all the users
- Call Dao to get list of users from map
- return error if map is empty
