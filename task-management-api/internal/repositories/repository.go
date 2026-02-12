package repositories

import (
	"gorm.io/gorm"
)

type GlobalRepository struct {
	User    UserRespositoryInterface
	Task    TaskRepositoryInterface
	Project ProjectRepositoryInterface
}

func NewGlobalRepository(DB *gorm.DB) *GlobalRepository {
	gr := &GlobalRepository{
		User:    &UserRepository{DB: DB},
		Task:    &TaskRepository{DB: DB},
		Project: &ProjectRepository{DB: DB},
	}
	return gr
}
