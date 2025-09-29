package postgres

import (
	"fmt"

	config "github.com/amenshenin/go_auth/internal/configs"
	"github.com/avast/retry-go"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func GetConnection(cfg *config.Config) (*sqlx.DB, error) {
	var db *sqlx.DB
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
			db, err = sqlx.Connect("pgx", dsn)
			return err
		},
		retry.Attempts(uint(cfg.DB.Retry.MaxAttempts)),
		retry.Delay(cfg.DB.Retry.Delay),
		retry.MaxDelay(cfg.DB.Retry.MaxDelay),
		retry.DelayType(retry.BackOffDelay),
	)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(cfg.DB.Pool.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.Pool.MaxIdleConns)
	return db, nil
}
