package main

import (
	"log"

	"github.com/Jiang-Gianni/DGAFstack/astra"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("keys.env")
	_ = astra.New()
	log.Println("Hello go")
	s := NewRESTServer(":4716")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
