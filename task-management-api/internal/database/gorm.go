package database

import (
	"fmt"

	"github.com/forester04/go-backend-projects/task-management-api/internal/errcode"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormClient() (*gorm.DB, error) {
	conString := viper.GetString("CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrConfigurationFailed, err)
	}

	err = migrate(db)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabaseMigrate, err)
	}

	return db, nil
}
