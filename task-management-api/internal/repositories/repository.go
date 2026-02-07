package repositories

import (
	"gorm.io/gorm"
)

type GlobalRepository struct {
	User UserRespositoryInterface
}

func NewGlobalRepository(DB *gorm.DB) *GlobalRepository {
	gr := &GlobalRepository{
		User: &UserRepository{DB: DB},
	}
	return gr
}
