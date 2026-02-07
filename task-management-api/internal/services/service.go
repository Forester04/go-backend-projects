package services

import (
	"github.com/forester04/go-backend-projects/task-management-api/internal/dto"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"github.com/forester04/go-backend-projects/task-management-api/internal/repositories"
	"go.uber.org/zap"
)

type ServiceInterface interface {
	RegisterUser(registerUser *dto.RegisterUser) (*models.User, error)
	LoginUser(email string, password string) (*models.User, error)
	formatRegisterUser(registerUser *dto.RegisterUser) (*models.User, error)
	DeleteUser(id uint) error
}

type Service struct {
	logger           *zap.Logger
	globalRepository *repositories.GlobalRepository
}

func New(logger *zap.Logger, globalRepossitory *repositories.GlobalRepository) *Service {
	service := &Service{
		logger,
		globalRepossitory,
	}
	return service
}
