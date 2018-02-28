package api

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type SliceStorage struct {
	userMu *sync.RWMutex
	users  []User
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{
		userMu: &sync.RWMutex{},
		users:  []User{},
	}
}

func (ss *SliceStorage) AddUser(user User) (User, error) {
	ss.userMu.Lock()
	defer ss.userMu.Unlock()

	user.ID = uuid.NewV4().String()
	ss.users = append(ss.users, user)
	return user, nil
}

func (ss *SliceStorage) GetUser(id string) (User, error) {
	ss.userMu.RLock()
	defer ss.userMu.RUnlock()

	for _, user := range ss.users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("No user with id: %s", id)
}

func (ss *SliceStorage) DeleteUser(id string) error {
	ss.userMu.Lock()
	defer ss.userMu.Unlock()

	for i, user := range ss.users {
		if user.ID == id {
			ss.users = append(ss.users[:i], ss.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No user with id: %s", id)
}
