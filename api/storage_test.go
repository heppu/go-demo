package api_test

import (
	"testing"

	"github.com/heppu/go-demo/api"
)

func BenchmarkSliceStorageSize100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := api.NewSliceStorage()
		for i := 0; i < 100; i++ {
			u, _ := s.AddUser(api.User{Name: "Test"})
			s.GetUser(u.ID)
		}
	}
}

func BenchmarkServerWithMapSize100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := api.NewMapStorage()
		for i := 0; i < 100; i++ {
			u, _ := s.AddUser(api.User{Name: "Test"})
			s.GetUser(u.ID)
		}
	}
}

func BenchmarkSliceStorageSize10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := api.NewSliceStorage()
		for i := 0; i < 10000; i++ {
			u, _ := s.AddUser(api.User{Name: "Test"})
			s.GetUser(u.ID)
		}
	}
}

func BenchmarkServerWithMapSize10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := api.NewMapStorage()
		for i := 0; i < 10000; i++ {
			u, _ := s.AddUser(api.User{Name: "Test"})
			s.GetUser(u.ID)
		}
	}
}
