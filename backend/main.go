package main

import (
	"log"

	"github.com/Jiang-Gianni/DGAFstack/astra"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("keys.env")
	astraDb := astra.New()
	log.Println(astraDb)
	// s := NewRESTServer(":4716")
	// if err := s.Run(); err != nil {
	// 	log.Fatal(err)
	// }
	//
}
