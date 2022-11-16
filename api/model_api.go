package api

import (
	userApi "refactoring"
	"refactoring/pkg/handler/requests"
)

type UserApiInterface interface {
	SearchUsers() (userApi.UserStore, error)
	CreateUser(DisplayName, Email string) (string, error)
	GetUser(id string) (userApi.User, error)
	UpdateUser(id string, request requests.UpdateUserRequest) error
	DeleteUser(id string) error
}
