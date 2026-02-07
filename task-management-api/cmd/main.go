package main

import (
	"fmt"
	"os"

	"github.com/forester04/go-backend-projects/task-management-api/internal/database"
	"github.com/forester04/go-backend-projects/task-management-api/internal/logger"
	"github.com/forester04/go-backend-projects/task-management-api/internal/repositories"
	"github.com/forester04/go-backend-projects/task-management-api/internal/services"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	initViper()

	//Initializing Logger
	logger, err := logger.New()
	if err != nil {
		panic(fmt.Errorf("error logger: %w", err))
	}
	defer func(Logger *zap.Logger) {
		//Sync logger before exit
		if err := Logger.Sync(); err != nil {
			logger.Fatal("error syncing logger", zap.Error(err))
		}
	}(logger)

	// Initializing Database
	gormClient, err := database.NewGormClient()
	if err != nil {
		logger.Fatal("error initializing database", zap.Error(err))
	}

	// Initializing Repository
	globalRepository := repositories.NewGlobalRepository(gormClient)

	//Initialize Service
	service := services.New(logger, globalRepository)

	//Initialize handlers

}

func initViper() {
	//Load environment variables
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error config file: %w", err))
	}
	if viper.GetString("HTTP_PROXY") != "" {
		os.Setenv("HTTP_PROXY", viper.GetString("HTTP_PROXY"))
	}
	if viper.GetString("HTTPS_PROXY") != "" {
		os.Setenv("HTTPS_PROXY", viper.GetString("HTTPS_PROXY"))
	}
}
