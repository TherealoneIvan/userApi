package api

import (
	"fmt"
	userApi "refactoring"
	"refactoring/pkg/handler/requests"
	"refactoring/pkg/repository"
	"strconv"
	"time"
)

const (
	userNotFoundErr  = "user not found"
	undefinedUserErr = "undefined user"
)

type UserApi struct {
	repository repository.UserApi
}

func NewUserApiService(repository repository.UserApi) *UserApi {
	return &UserApi{repository: repository}
}
func (u *UserApi) SearchUsers() (userApi.UserStore, error) {
	UserStore, err := u.repository.SearchUsers()
	if err != nil {
		return userApi.UserStore{}, err
	}
	return UserStore, nil
}
func (u *UserApi) CreateUser(DisplayName, Email string) (string, error) {
	UserStore, err := u.repository.SearchUsers()
	if err != nil {
		return "", err
	}
	UserStore.Increment++
	user := userApi.User{
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
	userStore, err := u.repository.SearchUsers()
	if err != nil {
		return err
	}

	if _, ok := userStore.List[id]; !ok {
		err := fmt.Errorf("%s", undefinedUserErr)
		return err
	}
	user := userStore.List[id]
	user.DisplayName = updateUserRequest.DisplayName
	userStore.List[id] = user
	u.repository.WriteChanges(userStore)
	return nil
}

func (u *UserApi) DeleteUser(id string) error {
	userStore, err := u.repository.DeleteUser()
	if err != nil {
		return err
	}
	if _, ok := userStore.List[id]; !ok {

		err := fmt.Errorf("%s", userNotFoundErr)
		return err
	}
	delete(userStore.List, id)
	u.repository.WriteChanges(userStore)
	return nil
}
func (u *UserApi) GetUser(id string) (userApi.User, error) {
	userStore, err := u.repository.SearchUsers()
	if err != nil {
		return userApi.User{}, err
	}
	return userStore.List[id], nil
}
