package repository

import (
	"context"
	"database/sql"
	"go-api-starter/model/entity"
)

type ExampleModelRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []entity.ExampleModel
	FindById(ctx context.Context, tx *sql.Tx, productId int) (entity.ExampleModel, error)
	Save(ctx context.Context, tx *sql.Tx, product entity.ExampleModel) entity.ExampleModel
}
