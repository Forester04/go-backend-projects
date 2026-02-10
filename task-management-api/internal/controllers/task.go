package controllers

import (
	"net/http"

	"github.com/forester04/go-backend-projects/task-management-api/internal/models"
	"github.com/forester04/go-backend-projects/task-management-api/internal/services"
	"github.com/forester04/go-backend-projects/task-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
)

func registerTaskRoutes(group *gin.RouterGroup, svc services.ServiceInterface) {
	group.POST("/", requestViewmodelMiddleware(&viewmodel.CreateTaskRequest{}), createTaskController(svc))
	group.GET("/:id", requestViewmodelMiddleware(&viewmodel.GetTaskRequest{}), getTaskController(svc))
	group.DELETE("/:id", requestViewmodelMiddleware(&viewmodel.DeleteTaskRequest{}), deleteTaskController(svc))
}

func createTaskController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.CreateTaskRequest)
		response := &viewmodel.CreateTaskResponse{}

		task := &models.Task{
			Title:       request.Body.Title,
			Description: request.Body.Description,
		}

		err := svc.CreateTask(task)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.ID = task.ID
		response.Body.Title = task.Title

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)
	}
}

func getTaskController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.GetTaskRequest)
		response := &viewmodel.GetTaskResponse{}

		task, err := svc.GetTask(request.ID)
		if err != nil {
			ctx.Error(err)
		}

		response.Body.ID = task.ID
		response.Body.Title = task.Title

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)
	}
}

func deleteTaskController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.DeleteProjectRequest)
		response := &viewmodel.DeleteProjectResponse{}

		err := svc.DeleteTask(request.ID)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.ID = request.ID

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseViewmodel, response)
	}
}
