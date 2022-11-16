package api

import (
	"fmt"
	userStore "refactoring"
	"refactoring/pkg/handler/requests"
	"refactoring/pkg/repository"
	"strconv"
	"time"
)

type UserApi struct {
	repository repository.UserApi
}

func NewUserApiService(repository repository.UserApi) *UserApi {
	return &UserApi{repository: repository}
}
func (u *UserApi) SearchUsers() (userStore.UserStore, error) {
	UserStore, err := u.repository.SearchUsers()
	if err != nil {
		return userStore.UserStore{}, err
	}
	return UserStore, nil
}
func (u *UserApi) CreateUser(DisplayName, Email string) (string, error) {
	UserStore, err := u.repository.SearchUsers()
	if err != nil {
		return "", err
	}
	UserStore.Increment++
	user := userStore.User{
		CreatedAt:   time.Now(),
		DisplayName: DisplayName,
		Email:       Email,
	}
	id := strconv.Itoa(UserStore.Increment)
	UserStore.List[id] = user
	u.repository.WriteChanges(UserStore)
	return id, nil
}
func (u *UserApi) UpdateUser(id string, updateUserRequest requests.UpdateUserRequest) error {
	UserStore, err := u.repository.SearchUsers()
	if err != nil {
		return err
	}

	if _, ok := UserStore.List[id]; !ok {
		err := fmt.Errorf("%s", "undefined user")
		return err
	}

	user := UserStore.List[id]
	user.DisplayName = updateUserRequest.DisplayName
	UserStore.List[id] = user
	u.repository.WriteChanges(UserStore)
	return nil
}

func (u *UserApi) DeleteUser(id string) error {
	UserStore, err := u.repository.DeleteUser()
	if err != nil {
		return err
	}
	if _, ok := UserStore.List[id]; !ok {
		err := fmt.Errorf("%s", "user not found")
		return err
	}
	delete(UserStore.List, id)
	u.repository.WriteChanges(UserStore)
	return nil
}
func (u *UserApi) GetUser(id string) (userStore.User, error) {
	UserStore, err := u.repository.SearchUsers()
	if err != nil {
		return userStore.User{}, err
	}
	return UserStore.List[id], nil
}
