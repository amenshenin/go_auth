package service

import (
	config "github.com/amenshenin/go_auth/internal/configs"
	"github.com/amenshenin/go_auth/internal/repository"
	"github.com/amenshenin/go_auth/internal/schemas"
)

const (
	StatusActive   = "A"
	StatusDisabled = "D"
	StatusHidden   = "H"
)

type InitCore interface {
	CreateTablesStructure() error
	CreateRootAdmin() error
}

type Autorization interface {
	CreateUser(user *schemas.UserInput) (int, error)
	GenerateToken(user *schemas.UserInput) (string, error)
	ParceToken(accessToken string) (int, error)
}

type Service struct {
	InitCore
	Autorization
}

func NewService(cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{
		InitCore:     NewCore(cfg, repo),
		Autorization: NewAuth(cfg, repo),
		// TodoList:     NewTodoListPostges(db),
	}
}
