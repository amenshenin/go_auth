package postgres

import (
	"fmt"

	config "github.com/amenshenin/go_auth/internal/configs"
	"github.com/amenshenin/go_auth/internal/storages"
	"github.com/avast/retry-go"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func GetConnection(cfg *config.Config) (*storages.Storage, error) {
	var db *sqlx.DB
	s := storages.Storage{DB: db}
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.DBName,
		cfg.DB.Password,
		cfg.DB.SSLMode,
	)
	err := retry.Do(
		func() error {
			var err error
			s.DB, err = sqlx.Connect("pgx", dsn)
			return err
		},
		retry.Attempts(uint(cfg.DB.Retry.MaxAttempts)),
		retry.Delay(cfg.DB.Retry.Delay),
		retry.MaxDelay(cfg.DB.Retry.MaxDelay),
		retry.DelayType(retry.BackOffDelay),
	)
	if err != nil {
		return &s, err
	}
	err = s.DB.Ping()
	if err != nil {
		return &s, err
	}
	s.DB.SetMaxOpenConns(cfg.DB.Pool.MaxOpenConns)
	s.DB.SetMaxIdleConns(cfg.DB.Pool.MaxIdleConns)
	return &s, nil
}
