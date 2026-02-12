package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/forester04/go-backend-projects/task-management-api/internal/errcode"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"gorm.io/gorm"
)

func (svc *Service) CreateProject(userID uint, project *models.Project) error {
	if err := svc.validateProject(project); err != nil {
		return err
	}

	exists, err := svc.globalRepository.Project.ExistsByName(userID, project.Name)
	if err != nil {
		return fmt.Errorf("%w: failed to check duplicates: %v", errcode.ErrDatabase, err)
	}
	if exists {
		return errcode.ErrDuplicate
	}

	err = svc.globalRepository.Project.Create(userID, project)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return nil
}

func (svc *Service) GetByProjectID(userID uint, id uint) (*models.Project, error) {
	project, err := svc.globalRepository.Project.GetByID(userID, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return project, nil
}

func (svc *Service) GetAllProjects(userID uint) ([]*models.Project, error) {
	projects, err := svc.globalRepository.Project.ListByUser(userID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return projects, nil
}

func (svc *Service) RenameProject(name string, userID, id uint) error {
	err := svc.globalRepository.Project.UpdateName(userID, id, name)
	if err != nil {
		return errcode.ErrDatabase
	}
	return nil
}

func (svc *Service) DeleteProject(userID, id uint) error {
	err := svc.globalRepository.Project.Delete(userID, id)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return nil
}

func (svc *Service) validateProject(project *models.Project) error {
	name := strings.TrimSpace(project.Name)

	if name == "" {
		return errcode.ErrInvalidParameters
	}

	if len(name) > 100 {
		return fmt.Errorf("%w: project name too long", errcode.ErrInvalidParameters)
	}

	if len(name) < 3 {
		return fmt.Errorf("%w: project name too long", errcode.ErrInvalidParameters)
	}
	project.Name = name
	return nil
}
