package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

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
	Crud CrudService[T],
	Sort func([]T, int, int) bool,
) error {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:
		typeContainers, err := Crud.Get()
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
		ids, err := Crud.Create([]T{*typeContainer})
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
	Crud CrudService[T],
) error {
	defer r.Body.Close()
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("invalid id")
	}
	switch r.Method {
	case http.MethodGet:
		typeContainers, err := Crud.GetById(id)
		if err != nil {
			return writeError(w, err)
		}
		return writeJSON(w, http.StatusOK, typeContainers)
	case http.MethodPut:
		toBeUpdated := new(T)
		json.NewDecoder(r.Body).Decode(toBeUpdated)
		err := Crud.Update([]T{*toBeUpdated})
		if err != nil {
			return writeError(w, err)
		}
		return writeJSON(w, http.StatusOK, id)
	case http.MethodDelete:
		err := Crud.Delete(id)
		if err != nil {
			return writeError(w, err)
		}
		return writeJSON(w, http.StatusOK, id)
	default:
		return writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func RegisterGenericCrud[T any](router *mux.Router, apiString string, crud CrudService[T], Sort func([]T, int, int) bool) {
	ctx := context.Background()
	router.HandleFunc(
		apiString,
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("HELLO???")
			if err := handleGeneric(ctx, w, r, crud, Sort); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
			}
		},
	)
	router.HandleFunc(
		apiString+"/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("HELLO???")
			if err := handleGenericById(ctx, w, r, crud); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
			}
		},
	)
}
