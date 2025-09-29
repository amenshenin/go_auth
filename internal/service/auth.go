package service

import "github.com/amenshenin/go_auth/internal/repository"

type Auth struct {
	r *repository.Repository
}

func NewAuth(r *repository.Repository) *Auth {
	return &Auth{r: r}
}
