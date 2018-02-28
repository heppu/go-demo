package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"testing"

	"github.com/heppu/go-demo/api"
)

func init() {
	go api.NewServer(":8000", api.NewMapStorage()).Start()
}

func TestServer(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			user := addUser(t)
			userExists(t, user)
			deleteUser(t, user)
			userNotExists(t, user)
		}()
	}

	wg.Wait()
}

func addUser(t *testing.T) api.User {
	resp, err := http.Post("http://localhost:8000/users", "application/json", bytes.NewBufferString(`{"Name": "Henri"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	user := api.User{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal(err)
	}
	return user
}

func userExists(t *testing.T, user api.User) {
	resp, err := http.Get("http://localhost:8000/users/" + user.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal(err)
	}
}

func userNotExists(t *testing.T, user api.User) {
	resp, err := http.Get("http://localhost:8000/users/" + user.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatal(err)
	}
}

func deleteUser(t *testing.T, user api.User) {
	req, err := http.NewRequest("DELETE", "http://localhost:8000/users/"+user.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
}
