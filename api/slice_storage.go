package api

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type SliceStorage struct {
	users []User
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{
		users: []User{},
	}
}

func (ss *SliceStorage) AddUser(user User) (User, error) {
	user.ID = uuid.NewV4().String()
	ss.users = append(ss.users, user)
	return user, nil
}

func (ss *SliceStorage) GetUser(id string) (User, error) {
	for _, user := range ss.users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("No user with id: %s", id)
}

func (ss *SliceStorage) DeleteUser(id string) error {
	for i, user := range ss.users {
		if user.ID == id {
			ss.users = append(ss.users[:i], ss.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No user with id: %s", id)
}
