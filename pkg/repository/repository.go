package repository

type Authtorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authtorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
