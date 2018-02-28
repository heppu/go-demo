package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"testing"

	"github.com/heppu/go-demo/api"
)

const (
	slicePort = "8000"
	mapPort   = "9000"
)

func init() {
	go api.NewServer(":"+slicePort, api.NewSliceStorage()).Start()
	go api.NewServer(":"+mapPort, api.NewMapStorage()).Start()
}

func TestServerWithSlice(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addAndDeleteUser(t, slicePort, wg)
	}
	wg.Wait()
}

func TestServeMapStorage(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addAndDeleteUser(t, mapPort, wg)
	}
	wg.Wait()
}

func addAndDeleteUser(t *testing.T, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	user := addUser(t, port)
	userExists(t, port, user)
	deleteUser(t, port, user)
	userNotExists(t, port, user)
}

func addUser(t *testing.T, port string) api.User {
	resp, err := http.Post("http://localhost:"+port+"/users", "application/json", bytes.NewBufferString(`{"Name": "Henri"}`))
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

func userExists(t *testing.T, port string, user api.User) {
	resp, err := http.Get("http://localhost:" + port + "/users/" + user.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal(err)
	}
}

func userNotExists(t *testing.T, port string, user api.User) {
	resp, err := http.Get("http://localhost:" + port + "/users/" + user.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatal(err)
	}
}

func deleteUser(t *testing.T, port string, user api.User) {
	req, err := http.NewRequest("DELETE", "http://localhost:"+port+"/users/"+user.ID, nil)
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
