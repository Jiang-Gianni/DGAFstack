package main

import (
	"net/http"

	"github.com/Jiang-Gianni/DGAFstack/astra"
	"github.com/Jiang-Gianni/DGAFstack/rest/mytable"
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
	RegisterGenericCrud[mytable.MyTable](
		router,
		"/api/v1/table",
		s.astraDb.MyTable(),
		func(t []mytable.MyTable, i, j int) bool {
			return t[i].Uuid < t[j].Uuid
		},
	)
	return http.ListenAndServe(s.listenAddress, router)
}
