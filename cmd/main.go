package main

import (
	"flag"
	"fmt"
	"log"

	config "github.com/amenshenin/go_auth/internal/configs"
	// "log/slog"
	// "os"
)

func main() {
	//Init config
	configPath := flag.String("config-path", "/run/secrets/config", "Please set in the config-path flag the path to config file")
	flag.Parse()
	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err.Error())
	}

	fmt.Printf("Config: %#v\n", config)

	//Init logs
	// log := initLogger(config.Enviremant)
	// log.Info("Start service: init logger complete")

	// //Init DB

	// //Init server

	// log.Info("go-go-go", config) //https://www.youtube.com/watch?v=rCJvW2xgnk0
}

// func initLogger(env string) *slog.Logger {
// 	var logg *slog.Logger
// 	switch env {
// 	case config.EnvLocal:
// 		logg = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
// 	case config.EnvDev:
// 		logg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
// 	case config.EnvProd:
// 		logg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
// 	default:
// 		log.Fatalf("Wrong log initialization")
// 	}
// 	return logg
// }
