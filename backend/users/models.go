package users

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserRequestBody struct {
	Name string `json:"name"`
}

type ResponseSingleUser struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type ResponseMultipleUsers struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Users    []User `json:"user"`
}
