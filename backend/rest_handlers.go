package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/Jiang-Gianni/DGAFstack/rest/user"
	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, err error) error {
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(err)
}

func handleGeneric[T any](
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	GetAll func() ([]T, error),
	Create func([]T) ([]string, error),
	Sort func([]T, int, int) bool,
) error {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:
		typeContainers, err := GetAll()
		if err != nil {
			return writeError(w, err)
		}
		sort.Slice(typeContainers[:], func(i, j int) bool {
			return Sort(typeContainers[:], i, j)
		})
		return writeJSON(w, http.StatusOK, typeContainers)
	case http.MethodPost:
		typeContainer := new(T)
		json.NewDecoder(r.Body).Decode(typeContainer)
		ids, err := Create([]T{*typeContainer})
		if err != nil {
			return writeError(w, err)
		}
		return writeJSON(w, http.StatusOK, ids)
	default:
		return writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func handleGenericById[T any](
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	GetById func(string) ([]T, error),
	Update func([]T) error,
	Delete func(string) error,
) error {
	defer r.Body.Close()
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("invalid id")
	}
	switch r.Method {
	case http.MethodGet:
		typeContainers, err := GetById(id)
		if err != nil {
			return writeError(w, err)
		}
		return writeJSON(w, http.StatusOK, typeContainers)
	case http.MethodPut:
		toBeUpdated := new(T)
		json.NewDecoder(r.Body).Decode(toBeUpdated)
		err := Update([]T{*toBeUpdated})
		if err != nil {
			return writeError(w, err)
		}
		return writeJSON(w, http.StatusOK, id)
	case http.MethodDelete:
		err := Delete(id)
		if err != nil {
			return writeError(w, err)
		}
		return writeJSON(w, http.StatusOK, id)
	default:
		return writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (s *RESTServer) handleUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return handleGeneric(
		ctx,
		w,
		r,
		s.astraDb.GetUsers,
		s.astraDb.CreateUsers,
		func(users []user.User, i, j int) bool {
			return users[i].Id < users[j].Id
		},
	)
}

func (s *RESTServer) handleUserById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return handleGenericById(
		ctx,
		w,
		r,
		s.astraDb.GetUserById,
		s.astraDb.UpdateUsers,
		s.astraDb.DeleteUserById,
	)
}
