package users

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(us *UserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var reqBody UserRequestBody

	json.NewDecoder(r.Body).Decode(&reqBody)

	user, err := uh.userService.CreateNewUser(r.Context(), reqBody.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp  := ResponseSingleUser{
		Error:   false,
		Message: "User created successfully",
		User:    *user,
	}

	json.NewEncoder(w).Encode(resp)
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users *[]User

	users, err := uh.userService.GetUsers(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := ResponseMultipleUsers{
		Error:   false,
		Message: "",
		Users:    *users,
	}

	json.NewEncoder(w).Encode(resp)

}
