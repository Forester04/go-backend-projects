package repositories

import (
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"gorm.io/gorm"
)

type TaskRepositoryInterface interface {
	Create(userID uint, projectID uint, task *models.Task) error
	GetByID(userID uint, projectID uint, id uint) (*models.Task, error)
	GetByTitle(userID uint, projectID uint, title string) (*models.Task, error)
	ListByUserProject(userID uint, projectID uint) ([]*models.Task, error)
	Save(userID uint, projectID uint, task *models.Task) error
	Delete(userID uint, projectID uint, id uint) error
}

type TaskRepository struct {
	DB *gorm.DB
}

func (rpt *TaskRepository) Create(userID uint, projectID uint, task *models.Task) error {
	return rpt.DB.Where("user_id = ? AND project_id = ?", userID, projectID).Create(task).Error
}

func (rpt *TaskRepository) GetByID(userID uint, projectID uint, id uint) (*models.Task, error) {
	task := &models.Task{}
	err := rpt.DB.Where("user_id = ? AND project_id = ?", userID, projectID).First(&task, id).Error
	return task, err
}

func (rpt *TaskRepository) GetByTitle(userID uint, projectID uint, title string) (*models.Task, error) {
	task := &models.Task{}
	err := rpt.DB.Where("user_id = ? AND project_id = ? AND title = ?", userID, projectID, title).First(&task).Error
	return task, err
}

func (rpt *TaskRepository) ListByUserProject(userID uint, projectID uint) ([]*models.Task, error) {
	tasks := []*models.Task{}
	err := rpt.DB.Where("user_id = ? AND project_id = ?", userID, projectID).Find(&tasks).Error
	return tasks, err
}

func (rpt *TaskRepository) Save(userID uint, projectID uint, task *models.Task) error {
	return rpt.DB.Where("user_id = ? AND project_id = ?", userID, projectID).Save(task).Error
}

func (rpt *TaskRepository) Delete(userID uint, projectID uint, id uint) error {
	return rpt.DB.Where("user_id = ? AND project_id = ? AND id = ?", userID, projectID, id).Delete(&models.Task{}).Error
}
