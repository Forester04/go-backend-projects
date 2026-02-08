package services

import (
	"github.com/forester04/go-backend-projects/task-management-api/internal/dto"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"github.com/forester04/go-backend-projects/task-management-api/internal/repositories"
	"go.uber.org/zap"
)

type ServiceInterface interface {
	/* Auth */
	RegisterUser(registerUser *dto.RegisterUser) (*models.User, error)
	LoginUser(email string, password string) (*models.User, error)
	formatRegisterUser(registerUser *dto.RegisterUser) (*models.User, error)
	DeleteUser(id uint) error
	GenerateToken(user *models.User) (tokenString string, err error)

	/* Task */
	CreateTask(task *models.Task) error
	GetTask(id uint) (*models.Task, error)
	DeleteTask(id uint) error

	/* Project */
	CreateProject(project *models.Project) error
	GetProject(id uint) (*models.Project, error)
	DeleteProject(id uint) error
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
