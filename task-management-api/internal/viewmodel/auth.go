package viewmodel

import "github.com/forester04/go-backend-projects/task-management-api/internal/dto"

type RegisterUserRequest struct {
	Body dto.RegisterUser `json:"body" binding:"required"`
}

type RegisterUserResponse struct {
	Body struct {
		Token string `json:"token"`
	} `json:"body"`
}

type LoginUserRequest struct {
	Body struct {
		Email    string `json:"eamil" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=64"`
	} `json:"body" binding:"required"`
}

type LoginUserRespoonse struct {
	Body struct {
		Token string `json:"token"`
	} `json:"body"`
}
