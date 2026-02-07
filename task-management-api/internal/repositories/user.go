package repositories

import (
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"gorm.io/gorm"
)

type UserRespositoryInterface interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	UpdateColumns(user *gorm.Model) error
	Update(user *models.User) error
	Delete(id uint) error
}

type UserRepository struct {
	DB *gorm.DB
}

func (rpt *UserRepository) Create(user *models.User) error {
	return rpt.DB.Create(user).Error
}

func (rpt *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := rpt.DB.Where("email = ?", email).Limit(1).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (rpt *UserRepository) GetByID(id uint) (*models.User, error) {
	user := &models.User{}
	err := rpt.DB.Where("id = ? ", id).First(user).Error
	if err != nil {

		return nil, err
	}
	return user, nil
}

func (rpt *UserRepository) UpdateColumns(user *gorm.Model) error {
	return rpt.DB.Model(user).Updates(user).Error
}

func (rpt *UserRepository) Update(user *models.User) error {
	return rpt.DB.UpdateColumns(user).Error
}

func (rpt *UserRepository) Delete(id uint) error {
	return rpt.DB.Delete(&models.User{}, id).Error
}
