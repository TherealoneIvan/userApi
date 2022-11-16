package repository

import userapi "refactoring"

type UserApi interface {
	SearchUsers() (userapi.UserStore, error)
	CreateUser() (userapi.UserStore, error)
	GetUser() (userapi.UserStore, error)
	UpdateUser() (userapi.UserStore, error)
	DeleteUser() (userapi.UserStore, error)
	WriteChanges(UserStore userapi.UserStore) error
}

type Repository struct {
	UserApi
}

func NewRepository(store string) *Repository {
	return &Repository{UserApi: NewJsonRepository(store)}
}
