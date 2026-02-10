package services

import (
	"fmt"

	"github.com/forester04/go-backend-projects/task-management-api/internal/errcode"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
)

func (svc *Service) CreateProject(project *models.Project) error {
	err := svc.globalRepository.Project.Create(project)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return nil
}

func (svc *Service) GetProject(id uint) (*models.Project, error) {
	project, err := svc.globalRepository.Project.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	if project.ID == 0 {
		return nil, fmt.Errorf("%w: %v", errcode.ErrNotFound, err)
	}
	return project, nil
}

func (svc *Service) GetAllProjects() ([]*models.Project, error) {
	projects, err := svc.globalRepository.Project.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return projects, nil
}

func (svc *Service) Rename(name string, id uint) (*models.Project, error) {
	project, err := svc.globalRepository.Project.Rename(name, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return project, nil
}

func (svc *Service) DeleteProject(id uint) error {
	err := svc.globalRepository.Project.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return nil
}
