package services

import (
	"fmt"

	"github.com/forester04/go-backend-projects/task-management-api/internal/errcode"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
)

func (svc *Service) CreateTask(task *models.Task) error {
	err := svc.globalRepository.Task.Create(task)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	return nil
}

func (svc *Service) GetTask(id uint) (*models.Task, error) {
	task, err := svc.globalRepository.Task.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	if task.ID == 0 {
		return nil, fmt.Errorf("%w: %v", errcode.ErrNotFound, err)
	}

	return task, nil

}

func (svc *Service) DeleteTask(id uint) error {
	err := svc.globalRepository.Task.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return nil
}
