package repository

import (
	"github.com/amenshenin/go_auth/internal/schemas"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type InitCore interface {
	CreateTablesStructure() error
}

type Autorization interface {
	CreateUser(user *schemas.UserInput, password string) (int, error)
	GetUser(user *schemas.UserInput, password string) (*schemas.User, error)
}

type Repository struct {
	InitCore
	Autorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		InitCore:     NewCore(db),
		Autorization: NewAuth(db),
		// TodoList:     NewTodoListPostges(db),
	}
}
