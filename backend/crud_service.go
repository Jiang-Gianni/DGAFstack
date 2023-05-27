package main

type CrudService[T any] interface {
	Get() ([]T, error)
	GetById(string) ([]T, error)
	Create([]T) ([]string, error)
	Update([]T) error
	Delete(string) error
}
