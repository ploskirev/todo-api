package service

import "github.com/ploskirev/todo-api/pkg/repository"

type Authtorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authtorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
