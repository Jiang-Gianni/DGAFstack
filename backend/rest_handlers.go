package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jiang-Gianni/DGAFstack/rest/user"
)

func (s *RESTServer) handleGetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO service get user
	user := []user.User{
		{
			ID:   1,
			Name: "Hello World",
		},
		{
			ID:   2,
			Name: "Bye Bye World",
		},
	}
	// user := &user.User{
	// 	ID:   1,
	// 	Name: "Hello World",
	// }
	s.count++
	fmt.Println(s.count)
	return writeJSON(w, http.StatusOK, user)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
