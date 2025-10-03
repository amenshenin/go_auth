package service

import (
	config "github.com/amenshenin/go_auth/internal/configs"
	"github.com/amenshenin/go_auth/internal/repository"
	// "github.com/amenshenin/go_auth/internal/schemas"
	// "golang.org/x/crypto/bcrypt"
)

type Core struct {
	cfg *config.Config
	r   *repository.Repository
}

func NewCore(cfg *config.Config, r *repository.Repository) *Core {
	return &Core{
		cfg: cfg,
		r:   r,
	}
}

func (c *Core) CreateTablesStructure() error {
	return c.r.InitCore.CreateTablesStructure()
}

func (c *Core) CreateRootAdmin() error {
	//TODO: need create
	return nil
}
