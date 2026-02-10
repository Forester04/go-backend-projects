package repositories

import (
	"fmt"

	"github.com/forester04/go-backend-projects/task-management-api/internal/errcode"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"gorm.io/gorm"
)

type TaskRepositoryInterface interface {
	Create(task *models.Task) error
	GetByID(id uint) (*models.Task, error)
	GetByTitle(name string) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	Rename(title string, id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

type TaskRepository struct {
	DB *gorm.DB
}

func (rpt *TaskRepository) Create(task *models.Task) error {
	return rpt.DB.Create(task).Error
}

func (rpt *TaskRepository) GetByID(id uint) (*models.Task, error) {
	task := &models.Task{}
	err := rpt.DB.Where("id = ?", id).First(task).Error
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return task, nil
}

func (rpt *TaskRepository) GetByTitle(name string) (*models.Task, error) {
	task := &models.Task{}
	err := rpt.DB.Where("name = ?", name).Limit(1).Find(task)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return task, nil
}

func (rpt *TaskRepository) GetAll() ([]*models.Task, error) {
	tasks := []*models.Task{}
	err := rpt.DB.Find(&tasks)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return tasks, nil
}

func (rpt *TaskRepository) Rename(title string, id uint) (*models.Task, error) {
	task, err := rpt.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	task.Title = title

	return task, nil
}

func (rpt *TaskRepository) Update(task *models.Task) error {
	return rpt.DB.UpdateColumns(task).Error
}

func (rpt *TaskRepository) Delete(id uint) error {
	return rpt.DB.Delete(&models.Project{}, id).Error
}
