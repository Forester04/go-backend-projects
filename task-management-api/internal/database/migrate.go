package database

import (
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(all()...); err != nil {
		return err
	}
	return nil
}

func all() []interface{} {
	out := []interface{}{}
	for _, v := range allMap() {
		out = append(out, v)
	}
	return out
}

func allMap() map[string]interface{} {
	return map[string]interface{}{
		"User": models.User{},
	}
}
