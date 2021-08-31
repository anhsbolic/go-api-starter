package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-api-starter/exception"
	"go-api-starter/helper"
	"go-api-starter/model/entity"
	"go-api-starter/model/web"
	"go-api-starter/repository"
)

type ExampleModelServiceImpl struct {
	ExampleModelRepository   repository.ExampleModelRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewExampleModelService(
	exampleModelRepository repository.ExampleModelRepository,
	db *sql.DB,
	validate *validator.Validate,
) ExampleModelService {
	return &ExampleModelServiceImpl{
		ExampleModelRepository:   exampleModelRepository,
		DB:                  db,
		Validate:            validate,
	}
}

func (service *ExampleModelServiceImpl) FindAll(ctx context.Context) []web.ExampleModelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	exampleModels := service.ExampleModelRepository.FindAll(ctx, tx)

	return web.ToExampleModelResponses(exampleModels)
}

func (service *ExampleModelServiceImpl) FindById(ctx context.Context, exampleModelId int) web.ExampleModelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	exampleModel, err := service.ExampleModelRepository.FindById(ctx, tx, exampleModelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ToExampleModelResponse(exampleModel)
}

func (service *ExampleModelServiceImpl) Create(ctx context.Context, request web.ExampleModelCreateRequest) web.ExampleModelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// create new exampleModel
	exampleModel := entity.ExampleModel{
		Name: request.Name,
	}
	exampleModel = service.ExampleModelRepository.Save(ctx, tx, exampleModel)

	return web.ToExampleModelResponse(exampleModel)
}