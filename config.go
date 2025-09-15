package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ProjectName string `ENV:"COMPOSE_PROJECT_NAME env-required:"true"`
	Enviremant  string `ENV:"ENVIRONMENT" env-default:"local"`
	DB          DB
	HTTPServer  HTTPServer
}

func containsWord(slice []string, word string) bool {
	for _, item := range slice {
		if item == word {
			return true
		}
	}
	return false
}

func (c Config) ValidateConfig() error {
	if c.ProjectName == "" {
		return fmt.Errorf("Config error. Empty required field %s", "ProjectName")
	}
	environiments := []string{"local", "dev", "prod"}
	if c.Enviremant == "" || !containsWord(environiments, c.Enviremant) {
		return fmt.Errorf("Config error. Empty required field %s", "Enviremant")
	}
	err := c.DB.validateDBConfig()
	if err != nil {
		return err
	}
	err = c.HTTPServer.validateHTTPServerConfig()
	if err != nil {
		return err
	}
	return nil
}

type DB struct {
	Host     string `ENV:"-"`
	Port     string `ENV:"-" env-default:"5432"`
	Username string `ENV:"DB_USER" env-required:"true"`
	Password string `ENV:"DB_PASS" env-required:"true"`
	DBName   string `ENV:"DB_NAME" env-required:"true"`
	SSLMode  string `ENV:"SSLMode" env-default:"disable"`
}

func (c DB) validateDBConfig() error {
	v := reflect.ValueOf(c)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		if fieldValue == "" {
			return fmt.Errorf("Config error. Empty required field %s", fieldName)
		}
	}
	return nil
}

type HTTPServer struct {
	Address     string        `ENV:"VHOST" env-required:"true"`
	Port        string        `ENV:"VHOST_PORT" env-required:"true"`
	Timeout     time.Duration `ENV:"TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `ENV:"IDDLE_TIMEOUT" env-default:"60s"`
}

func (c HTTPServer) validateHTTPServerConfig() error {
	v := reflect.ValueOf(c)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		if fieldValue == "" {
			return fmt.Errorf("Config error. Empty required field %s", fieldName)
		}
	}
	return nil
}

func MustLoad(path string) *Config {
	var cfg Config
	if path == "" {
		log.Fatal("Config file is not set")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("%s. Config file %s does not exist", err.Error(), path)
	}
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("%s. Cannot read the config file %s", err.Error(), path)
	}
	t, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		log.Fatalf("Config error. Empty required field %s", "Timeout")
	}
	it, err := time.ParseDuration(os.Getenv("IDDLE_TIMEOUT"))
	if err != nil {
		log.Fatalf("Config error. Empty required field %s", "IdleTimeout")
	}
	cfg = Config{
		ProjectName: os.Getenv("COMPOSE_PROJECT_NAME"),
		Enviremant:  os.Getenv("ENVIRONMENT"),
		DB: DB{
			Host:     fmt.Sprintf("%s-db", cfg.ProjectName),
			Port:     "5432",
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("SSL_MODE"),
		},
		HTTPServer: HTTPServer{
			Address:     os.Getenv("VHOST"),
			Port:        os.Getenv("VHOST_PORT"),
			Timeout:     t,
			IdleTimeout: it,
		},
	}
	err = cfg.ValidateConfig()
	if err != nil {
		log.Fatalf("Config error. %s", err.Error())
	}

	return &cfg
}
