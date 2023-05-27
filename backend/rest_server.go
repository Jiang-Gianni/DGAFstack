package main

import (
	"context"
	"net/http"

	"github.com/Jiang-Gianni/DGAFstack/astra"
	"github.com/gorilla/mux"
)

type RESTServer struct {
	listenAddress string
	astraDb       *astra.AstraDB
}

func NewRESTServer(listenAddress string, astraDb *astra.AstraDB) *RESTServer {
	return &RESTServer{
		listenAddress: listenAddress,
		astraDb:       astraDb,
	}
}

func (s *RESTServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/user", adaptHttpHandlerFunc(s.handleUser))
	router.HandleFunc("/user/{id}", adaptHttpHandlerFunc(s.handleUserById))
	return http.ListenAndServe(s.listenAddress, nil)
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func adaptHttpHandlerFunc(apiFunc APIFunc) http.HandlerFunc {
	ctx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}
