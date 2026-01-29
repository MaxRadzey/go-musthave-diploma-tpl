package app

import (
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/config"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/logger"
	"go.uber.org/zap"
)

// Run запускает http сервер.
func Run(AppConfig *config.Config) error {
	if err := logger.Initialize(AppConfig.LogLevel); err != nil {
		return err
	}

	// storageResult, err := dbstorage.InitializeStorage(AppConfig.DatabaseDSN, AppConfig.FilePath)
	// if err != nil {
	// 	return err
	// }

	// urlService := service.NewService(storageResult.Storage, *AppConfig)
	// h := &httphandlers.Handler{Service: urlService}

	// r := router.SetupRouter(h, AppConfig)

	logger.Log.Info("Starting HTTP server", zap.String("address", AppConfig.RunAddress))
	// return r.Run(AppConfig.RunAddress)
	return nil
}
