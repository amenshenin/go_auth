package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Core struct {
	db *sqlx.DB
}

func NewCore(db *sqlx.DB) *Core {
	return &Core{db: db}
}

func (c *Core) CreateTablesStructure() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    id            serial       not null unique,
    name          varchar(255) not null unique,
    password_hash varchar(255) not null,
	status		  char(1) not null default 'D'
);`, usersTable)
	_, err := c.db.Exec(query)
	return err
}
