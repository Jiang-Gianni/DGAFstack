package main

import (
	"context"
	"net/http"
)

type RESTServer struct {
	listenAddress string
	count         int
}

func NewRESTServer(listenAddress string) *RESTServer {
	return &RESTServer{
		listenAddress: listenAddress,
		count:         0,
	}
}

func (s *RESTServer) Run() error {
	http.HandleFunc("/", adaptHttpHandlerFunc(s.handleGetUser))
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
