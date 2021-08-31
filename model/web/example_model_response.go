package web

import "go-api-starter/model/entity"

type ExampleModelResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToExampleModelResponse(product entity.ExampleModel) ExampleModelResponse {
	return ExampleModelResponse{
		Id:   product.Id,
		Name: product.Name,
	}
}

func ToExampleModelResponses(products []entity.ExampleModel) []ExampleModelResponse {
	var productResponses []ExampleModelResponse
	for _, product := range products {
		productResponses = append(productResponses, ToExampleModelResponse(product))
	}
	return productResponses
}
