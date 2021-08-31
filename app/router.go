package app

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"go-api-starter/controller"
	"go-api-starter/exception"
	"go-api-starter/repository"
	"go-api-starter/service"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *httprouter.Router {
	router := httprouter.New()

	// init repositories
	exampleModelRepository := repository.NewExampleModelRepository()

	// init service & controller
	exampleModelService := service.NewExampleModelService(exampleModelRepository, db, validate)
	exampleModelController := controller.NewExampleModelController(exampleModelService)

	// routes
	router.GET("/api/example-models", exampleModelController.FindAll)
	router.POST("/api/example-models", exampleModelController.Create)
	router.GET("/api/example-models/:exampleModelId", exampleModelController.FindById)

	// router handler
	router.PanicHandler = exception.ErrorHandler

	return router
}
