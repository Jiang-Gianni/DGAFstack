package main

import (
	"log"

	"github.com/Jiang-Gianni/DGAFstack/astra"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("keys.env")
	astraDb := astra.New()
	s := NewRESTServer(":4716", astraDb)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
