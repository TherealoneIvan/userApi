package repository

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	userapi "refactoring"
)

//const store = `users.json`

type RepositoryJSON struct {
	store string
}

func NewJsonRepository(store string) *RepositoryJSON {
	return &RepositoryJSON{store: store}
}
func (r *RepositoryJSON) userStore() (userapi.UserStore, error) {
	f, err := ioutil.ReadFile(r.store)
	if err != nil {
		return userapi.UserStore{}, err
	}
	s := userapi.UserStore{}
	err = json.Unmarshal(f, &s)
	if err != nil {
		return userapi.UserStore{}, err
	}
	return s, nil
}

func (r *RepositoryJSON) CreateUser() (userapi.UserStore, error) {
	return r.userStore()
}
func (r *RepositoryJSON) GetUser() (userapi.UserStore, error) {
	return r.userStore()
}
func (r *RepositoryJSON) UpdateUser() (userapi.UserStore, error) {
	return r.userStore()
}
func (r *RepositoryJSON) DeleteUser() (userapi.UserStore, error) {
	return r.userStore()
}

func (r *RepositoryJSON) SearchUsers() (userapi.UserStore, error) {
	return r.userStore()
}
func (r *RepositoryJSON) WriteChanges(UserStore userapi.UserStore) error {
	byteSlice, err := json.Marshal(&UserStore)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(r.store, byteSlice, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
