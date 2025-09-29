package repository

import "github.com/jmoiron/sqlx"

type Autorization interface {
	// CreateUser(user todo.User) (int, error)
	// GetUser(username, password string) (todo.User, error)
}

type Repository struct {
	Autorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuth(db),
		// TodoList:     NewTodoListPostges(db),
	}
}
