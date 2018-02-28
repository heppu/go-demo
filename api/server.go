package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	addr    string
	storage UserStorage
}

type UserStorage interface {
	AddUser(User) (User, error)
	GetUser(id string) (User, error)
	DeleteUser(id string) error
}

type User struct {
	ID   string
	Name string
}

func NewServer(addr string, storage UserStorage) *Server {
	return &Server{
		addr:    addr,
		storage: storage,
	}
}

func (s *Server) Start() error {
	handler := httprouter.New()
	handler.POST("/users", s.addUser)
	handler.GET("/users/:id", s.getUser)
	handler.DELETE("/users/:id", s.deleteUser)

	return http.ListenAndServe(s.addr, handler)
}

func (s *Server) addUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.storage.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := s.storage.GetUser(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if err := s.storage.DeleteUser(p.ByName("id")); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
