package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
	EnvDev   = "dev"
)

type Config struct {
	Enviremant string `env:"ENVIRONMENT" env-default:"local"`
	DB         struct {
		Host     string `env:"DB_HOST" env-required:"true"`
		Port     int    `env:"DB_PORT" env-default:"5432"`
		Username string `env:"DB_USER" env-required:"true"`
		Password string `env:"DB_PASS" env-required:"true"`
		DBName   string `env:"DB_NAME" env-required:"true"`
		SSLMode  string `env:"DB_SSLMode" env-default:"disable"`

		Pool struct {
			MaxOpenConns int `env:"DB_MAX_OPEN_CONNS" env-default:"10"`
			MaxIdleConns int `env:"DB_MAX_IDLE_CONNS" env-default:"5"`
		}

		Retry struct {
			MaxAttempts int           `env:"DB_MAX_ATTEMPTS" env-default:"10"`
			Delay       time.Duration `env:"DB_DELAY" env-default:"2s"`
			MaxDelay    time.Duration `env:"DB_MAX_DELAY" env-default:"10s"`
		}
	}
	HTTPServer struct {
		Address     string        `env:"VHOST" env-required:"true"`
		Port        string        `env:"VHOST_PORT" env-required:"true"`
		Timeout     time.Duration `env:"SERV_TIMEOUT" env-default:"4s"`
		IdleTimeout time.Duration `env:"SERV_IDDLE_TIMEOUT" env-default:"60s"`
	}
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	if path == "" {
		return &cfg, errors.New("Config file is not set")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &cfg, fmt.Errorf("%w. Config file %s does not exist", err, path)
	}
	err := godotenv.Load(path)
	if err != nil {
		return &cfg, fmt.Errorf("%w. Cannot read the config file %s", err, path)
	}
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &cfg, fmt.Errorf("%w. Cannot read from environment variables", err)
	}
	return &cfg, nil
}
