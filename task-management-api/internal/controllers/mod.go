package controllers

import (
	"reflect"
	"strings"

	"github.com/forester04/go-backend-projects/task-management-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Router struct {
	engine *gin.Engine
	logger *zap.Logger
}

func NewRouter(logger *zap.Logger, svc services.ServiceInterface) *Router {
	router := &Router{}

	router.logger = logger
	router.engine = gin.Default()
	config(router.engine)

	/* Middlewares */
	router.engine.Use(router.corsMiddleware())
	router.engine.Use(router.responseViewmodelMiddleware())
	router.engine.Use(router.errorHandlerMiddleware())

	router.registerRoutes(svc)

	return router

}

func (rtr *Router) Run(addr string) error {
	return rtr.engine.Run(addr)
}

func (rtr *Router) registerRoutes(svc services.ServiceInterface) {
	/* Auth */
	auth := rtr.engine.Group("/auth")
	registerUserRoutes(auth, svc)

	/* Task */
	task := rtr.engine.Group("/task")
	registerTaskRoutes(task, svc)

	/* Project */
	project := rtr.engine.Group("/project")
	registerProjectRoutes(project, svc)

}

func config(router *gin.Engine) {
	router.RedirectTrailingSlash = false

	// custom validator
	//This is used to be able to get the json tag name instead of the struct field name, to send to the client, the client can use it to display the error message
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return field.Name
			}
			return name
		})
	}

}
