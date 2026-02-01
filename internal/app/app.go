package app

import (
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/config"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/logger"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/storage"
	"go.uber.org/zap"
)

// Run запускает http сервер.
func Run(cfg *config.Config) error {
	if err := logger.Initialize(cfg.LogLevel); err != nil {
		return err
	}

	storageResult, err := storage.InitializeStorage(cfg.DatabaseDSN)
	if err != nil {
		return err
	}
	defer storageResult.Close()

	// userService := service.New(storageResult.UserRepository)
	// h := handler.New(userService)
	// r := router.SetupRouter(h, cfg)
	logger.Log.Info("Starting HTTP server", zap.String("address", cfg.RunAddress))
	// return r.Run(cfg.RunAddress)
	return nil
}
