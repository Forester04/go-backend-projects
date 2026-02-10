package controllers

import (
	"net/http"

	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"github.com/forester04/go-backend-projects/task-management-api/internal/services"
	"github.com/forester04/go-backend-projects/task-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
)

func registerProjectRoutes(group *gin.RouterGroup, svc services.ServiceInterface) {
	group.POST("/", requestViewmodelMiddleware(&viewmodel.CreateProjectRequest{}), createProjectController(svc))
	group.GET("/:id", requestViewmodelMiddleware(&viewmodel.GetProjectRequest{}), getProjectController(svc))
	group.DELETE("/:id", requestViewmodelMiddleware(&viewmodel.DeleteProjectRequest{}), deleteProjectController(svc))
}

func createProjectController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.CreateProjectRequest)
		response := &viewmodel.CreateProjectResponse{}

		project := &models.Project{
			Name:        request.Body.Name,
			Description: request.Body.Description,
		}

		err := svc.CreateProject(project)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.ID = project.ID
		response.Body.Name = project.Name

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)
	}
}

func getProjectController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.GetProjectRequest)
		response := &viewmodel.GetProjectResponse{}

		project, err := svc.GetProject(request.ID)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.ID = project.ID
		response.Body.Name = project.Name

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)

	}
}

func deleteProjectController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.DeleteProjectRequest)
		response := &viewmodel.DeleteProjectResponse{}

		err := svc.DeleteProject(request.ID)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.ID = request.ID

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)
	}
}
