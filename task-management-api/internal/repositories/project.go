package repositories

import (
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	Create(userID uint, project *models.Project) error
	GetByID(userID uint, projectID uint) (*models.Project, error)
	GetByName(userID uint, name string) (*models.Project, error)
	ListByUser(userID uint) ([]*models.Project, error)
	Save(userID uint, project *models.Project) error
	Delete(userID uint, id uint) error
}
type ProjectRepository struct {
	DB *gorm.DB
}

func (rpt *ProjectRepository) Create(userID uint, project *models.Project) error {
	return rpt.DB.Where("user_id = ?", userID).Create(project).Error
}

func (rpt *ProjectRepository) GetByID(userID uint, id uint) (*models.Project, error) {
	project := &models.Project{}
	err := rpt.DB.Where("user_id = ? AND  id = ?", userID, id).Find(&project).Error
	return project, err
}

func (rpt *ProjectRepository) ListByUser(userID uint) ([]*models.Project, error) {
	projects := []*models.Project{}
	err := rpt.DB.Where("user_id = ?", userID).Find(&projects).Error
	return projects, err
}

func (rpt *ProjectRepository) GetByName(userID uint, name string) (*models.Project, error) {
	project := &models.Project{}
	err := rpt.DB.Where("user_id = ? AND name = ?", userID, name).First(&project).Error
	return project, err
}

func (rpt *ProjectRepository) Save(userID uint, project *models.Project) error {
	return rpt.DB.Where("user_id = ?", userID).Save(project).Error
}

func (rpt *ProjectRepository) Delete(userID uint, id uint) error {
	return rpt.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Project{}).Error
}
