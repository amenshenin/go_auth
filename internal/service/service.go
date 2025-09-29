package service

import (
	"github.com/amenshenin/go_auth/internal/repository"
)

type Autorization interface {
	// CreateUser(user todo.User) (int, error)
	// GetUser(username, password string) (todo.User, error)
}

type Service struct {
	Autorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuth(repo),
		// TodoList:     NewTodoListPostges(db),
	}
}
