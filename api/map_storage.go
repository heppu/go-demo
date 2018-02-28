package api

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MapStorage struct {
	userMu *sync.RWMutex
	users  map[string]User
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		userMu: &sync.RWMutex{},
		users:  map[string]User{},
	}
}

func (ms *MapStorage) AddUser(user User) (User, error) {
	ms.userMu.Lock()
	defer ms.userMu.Unlock()

	user.ID = uuid.NewV4().String()
	ms.users[user.ID] = user
	return user, nil
}

func (ms *MapStorage) GetUser(id string) (User, error) {
	ms.userMu.RLock()
	defer ms.userMu.RUnlock()

	var err error
	user, ok := ms.users[id]
	if !ok {
		err = fmt.Errorf("No user with id: %s", id)
	}

	return user, err
}

func (ms *MapStorage) DeleteUser(id string) error {
	ms.userMu.Lock()
	defer ms.userMu.Unlock()

	if _, ok := ms.users[id]; !ok {
		return fmt.Errorf("No user with id: %s", id)
	}

	delete(ms.users, id)
	return nil
}
