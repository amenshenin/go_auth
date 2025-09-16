package loggers

import (
	"errors"
	"log/slog"
	"os"

	config "github.com/amenshenin/go_auth/internal/configs"
)

func GetLogger(cfg *config.Config) (*slog.Logger, error) {
	var logg *slog.Logger
	switch cfg.Enviremant {
	case config.EnvLocal:
		logg = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvDev:
		logg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProd:
		logg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		return logg, errors.New("wrong log initialization")
	}
	return logg, nil
}
