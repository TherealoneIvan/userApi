package api

import (
	userapi "refactoring"
	"refactoring/pkg/handler/requests"
)

type UserApiInterface interface {
	SearchUsers() (userapi.UserStore, error)
	CreateUser(DisplayName, Email string) (string, error)
	GetUser(id string) (userapi.User, error)
	UpdateUser(id string, request requests.UpdateUserRequest) error
	DeleteUser(id string) error
}
