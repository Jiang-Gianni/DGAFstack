package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello go")
	s := NewRESTServer(":4716")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
