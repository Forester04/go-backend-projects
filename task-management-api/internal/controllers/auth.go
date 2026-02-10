package controllers

import (
	"net/http"

	"github.com/forester04/go-backend-projects/task-management-api/internal/services"
	"github.com/forester04/go-backend-projects/task-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(group *gin.RouterGroup, svc services.ServiceInterface) {
	group.POST("/register", requestViewmodelMiddleware(&viewmodel.RegisterUserRequest{}), registerController(svc))
	group.POST("/login", requestViewmodelMiddleware(&viewmodel.LoginUserRequest{}), loginController(svc))

}
func registerController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.RegisterUserRequest)
		response := &viewmodel.RegisterUserResponse{}

		user, err := svc.RegisterUser(&request.Body)
		if err != nil {
			ctx.Error(err)
			return
		}

		token, err := svc.GenerateToken(user)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.Token = token

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)

	}
}

func loginController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.LoginUserRequest)
		response := &viewmodel.LoginUserRespoonse{}

		user, err := svc.LoginUser(request.Body.Email, request.Body.Password)
		if err != nil {
			ctx.Error(err)
			return
		}

		token, err := svc.GenerateToken(user)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.Token = token

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)
	}
}
