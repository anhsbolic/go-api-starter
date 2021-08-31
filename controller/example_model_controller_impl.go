package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-api-starter/helper"
	"go-api-starter/model/web"
	"go-api-starter/service"
	"net/http"
	"strconv"
)

type ExampleModelControllerImpl struct {
	ExampleModelService service.ExampleModelService
}

func NewExampleModelController(exampleModelService service.ExampleModelService) ExampleModelController {
	return &ExampleModelControllerImpl{ExampleModelService: exampleModelService}
}

func (controller *ExampleModelControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleModelResponses := controller.ExampleModelService.FindAll(request.Context())
	jsonResponse := web.JSONResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   exampleModelResponses,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, jsonResponse)
}

func (controller *ExampleModelControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleModelId := params.ByName("exampleModelId")
	id, err := strconv.Atoi(exampleModelId)
	helper.PanicIfError(err)

	exampleModelResponse := controller.ExampleModelService.FindById(request.Context(), id)
	jsonResponse := web.JSONResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   exampleModelResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, jsonResponse)
}

func (controller *ExampleModelControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleModelCreateRequest := web.ExampleModelCreateRequest{}
	helper.ReadFromRequestBody(request, &exampleModelCreateRequest)

	exampleModelResponse := controller.ExampleModelService.Create(request.Context(), exampleModelCreateRequest)
	jsonResponse := web.JSONResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   exampleModelResponse,
	}

	writer.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(writer, jsonResponse)
}
