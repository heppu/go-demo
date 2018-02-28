package main

import (
	"log"

	"github.com/heppu/go-demo/api"
)

func main() {
	if err := api.NewServer(":8000", api.NewSliceStorage()).Start(); err != nil {
		log.Fatal(err)
	}
}
