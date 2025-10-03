package repository

import (
	"fmt"

	"github.com/amenshenin/go_auth/internal/schemas"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user *schemas.UserInput, password string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash) VALUES ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth) GetUser(user *schemas.UserInput, password string) (*schemas.User, error) {
	var u schemas.User
	query := fmt.Sprintf("SELECT id, name, password_hash, status FROM %s WHERE name = $1 AND password_hash = $2 AND status = $3", usersTable)
	err := r.db.Get(&u, query, user.Name, password, user.Status)
	return &u, err
}
