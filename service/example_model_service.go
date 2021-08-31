package service

import (
	"context"
	"go-api-starter/model/web"
)

type ExampleModelService interface {
	FindAll(ctx context.Context) []web.ExampleModelResponse
	FindById(ctx context.Context, productId int) web.ExampleModelResponse
	Create(ctx context.Context, request web.ExampleModelCreateRequest) web.ExampleModelResponse
}
