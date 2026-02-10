package repositories

import (
	"fmt"

	"github.com/forester04/go-backend-projects/task-management-api/internal/errcode"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	Create(project *models.Project) error
	GetByID(id uint) (*models.Project, error)
	GetByName(name string) (*models.Project, error)
	GetAll() ([]*models.Project, error)
	Update(project *models.Project) error
	Rename(name string, id uint) (*models.Project, error)
	Delete(id uint) error
}
type ProjectRespository struct {
	DB *gorm.DB
}

func (rpt *ProjectRespository) Create(project *models.Project) error {
	return rpt.DB.Create(project).Error
}

func (rpt *ProjectRespository) GetByID(id uint) (*models.Project, error) {
	project := &models.Project{}
	err := rpt.DB.Where("id = ?", id).First(project)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return project, nil
}

func (rpt *ProjectRespository) GetAll() ([]*models.Project, error) {
	projects := []*models.Project{}
	err := rpt.DB.Find(&projects)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return projects, nil
}

func (rpt *ProjectRespository) Rename(name string, id uint) (*models.Project, error) {
	project, err := rpt.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	project.Name = name
	return project, nil
}

func (rpt *ProjectRespository) GetByName(name string) (*models.Project, error) {
	project := &models.Project{}
	err := rpt.DB.Where("name = ?", name).Limit(1).Find(project)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return project, nil
}

func (rpt *ProjectRespository) Update(project *models.Project) error {
	return rpt.DB.UpdateColumns(project).Error
}

func (rpt *ProjectRespository) Delete(id uint) error {
	return rpt.DB.Delete(&models.Project{}, id).Error
}
