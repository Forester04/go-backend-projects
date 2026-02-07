package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/forester04/go-backend-projects/task-management-api/internal/dto"
	"github.com/forester04/go-backend-projects/task-management-api/internal/errcode"
	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (svc *Service) RegisterUser(registerUser *dto.RegisterUser) (*models.User, error) {
	user, err := svc.globalRepository.User.GetByEmail(registerUser.Email)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("%w", errcode.ErrUserAlreadyExists)
	}

	user, err = svc.formatRegisterUser(registerUser)
	if err != nil {
		return nil, err
	}

	err = svc.globalRepository.User.Create(user)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	return user, nil

}

func (svc *Service) LoginUser(email string, password string) (*models.User, error) {
	user, err := svc.globalRepository.User.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	if user == nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrInvalidCredentials, errors.New("user does not exist"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrInvalidCredentials, err)
	}
	return user, nil
}

func (svc *Service) formatRegisterUser(registerUser *dto.RegisterUser) (*models.User, error) {
	registerUser.Email = strings.ToLower(strings.TrimSpace(registerUser.Email))

	registerUser.Username = strings.ToLower(registerUser.Username)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 12)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrExternalLib, err)
	}
	registerUser.Password = string(passwordHash)

	var parsedBirthDate time.Time
	if registerUser.BirthDate != "" {
		parsedBirthDate, err = time.Parse("2006-01-02", registerUser.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", errcode.ErrInvalidParameters, err)
		}
	}

	user := &models.User{
		Email:     registerUser.Email,
		Password:  registerUser.Password,
		Username:  registerUser.Username,
		FirstName: registerUser.FirstName,
		LastName:  registerUser.LastName,
		Phone:     &registerUser.Phone,
		BirthDate: &parsedBirthDate,
	}
	return user, nil
}

func (svc *Service) DeleteUser(id uint) error {
	user, err := svc.globalRepository.User.GetByID(id)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	if user.ID == 0 {
		return fmt.Errorf("%w: %v", errcode.ErrNotFound, err)
	}

	err = svc.globalRepository.User.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return nil
}
