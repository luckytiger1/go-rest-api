package main

import (
	"log"

	"github.com/luckytiger1/go-rest-api.git/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
