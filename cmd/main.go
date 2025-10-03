package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/amenshenin/go_auth/internal/configs"
	"github.com/amenshenin/go_auth/internal/handler"
	"github.com/amenshenin/go_auth/internal/httpserver"
	"github.com/amenshenin/go_auth/internal/loggers"
	"github.com/amenshenin/go_auth/internal/repository"
	"github.com/amenshenin/go_auth/internal/service"
	"github.com/amenshenin/go_auth/internal/storages/postgres"
	// "log/slog"
	// "os"
)

func main() {
	//Init config
	configPath := flag.String("config-path", "/run/secrets/config", "Please set in the config-path flag the path to config file")
	needCreteStructure := flag.Bool("need-create-structure", false, "Please set the need-create-structure if you need create tables and primary data")
	flag.Parse()
	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err.Error())
	}

	//Init logs
	logger, err := loggers.GetLogger(config)
	if err != nil {
		log.Fatalf("Error init logger: %s", err.Error())
	}
	logger.Info("Start service: init logger complete")

	//Init DB
	db, err := postgres.GetConnection(config)
	if err != nil {
		logger.Error("Error database connection", "error", err.Error())
		os.Exit(1)
	}
	logger.Info("Start service: getting database connection complete")

	repo := repository.NewRepository(db)
	service := service.NewService(config, repo)
	if *needCreteStructure {
		err := service.InitCore.CreateTablesStructure()
		if err != nil {
			logger.Error("Cannot create tables structure", "error", err.Error())
			os.Exit(1)
		}
		logger.Info("Tables structure has been created successfully")
		err = service.InitCore.CreateRootAdmin()
		if err != nil {
			logger.Error("Cannot create root admin", "error", err.Error())
			os.Exit(1)
		}
		logger.Info("Root admin has been created successfully")
	}
	handlers := handler.NewHandler(config, logger, service)
	server := new(httpserver.Server)

	go func() {
		if err := server.Run(config, handlers.InitRouts()); err != nil {
			logger.Error("error occured while running http server", "error", err.Error())
			os.Exit(1)
		}
	}()
	logger.Info("server started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done
	logger.Info("stopping server")

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error("error occured on server shutting down", "error", err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Error("error occured on db connection close", "error", err.Error())
	}
}
