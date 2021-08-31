package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-api-starter/helper"
	"go-api-starter/model/entity"
)

type ExampleModelRepositoryImpl struct {
}

func NewExampleModelRepository() ExampleModelRepository {
	return &ExampleModelRepositoryImpl{}
}

func (repository *ExampleModelRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.ExampleModel {
	SQL := "select id, name from example_models"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var exampleModels []entity.ExampleModel
	for rows.Next() {
		exampleModel := entity.ExampleModel{}
		err := rows.Scan(&exampleModel.Id, &exampleModel.Name)
		helper.PanicIfError(err)
		exampleModels = append(exampleModels, exampleModel)
	}

	return exampleModels
}

func (repository *ExampleModelRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, exampleModelId int) (entity.ExampleModel, error) {
	SQL := "select id, name from example_models where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, exampleModelId)
	helper.PanicIfError(err)
	defer rows.Close()

	exampleModel := entity.ExampleModel{}
	if rows.Next() {
		err := rows.Scan(&exampleModel.Id, &exampleModel.Name)
		helper.PanicIfError(err)

		return exampleModel, nil
	} else {
		return exampleModel, errors.New("example model not found")
	}
}

func (repository *ExampleModelRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, exampleModel entity.ExampleModel) entity.ExampleModel {
	SQL := "insert into example_models(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, exampleModel.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	exampleModel.Id = int(id)
	return exampleModel
}
