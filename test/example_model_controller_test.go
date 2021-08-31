package test

import (
	"context"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"go-api-starter/model/entity"
	"go-api-starter/repository"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetExampleModelsSuccess(t *testing.T) {
	db := SetupTestDB()
	ClearExampleModelsTable(db)

	tx, _ := db.Begin()
	exampleModelRepository := repository.NewExampleModelRepository()
	exampleModel1 := exampleModelRepository.Save(context.Background(), tx, entity.ExampleModel{
		Name: "Something",
	})
	exampleModel2 := exampleModelRepository.Save(context.Background(), tx, entity.ExampleModel{
		Name: "Something 2",
	})
	tx.Commit()

	router := SetupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/example-models", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	var exampleModels = responseBody["data"].([]interface{})
	exampleModelResponse1 := exampleModels[0].(map[string]interface{})
	exampleModelResponse2 := exampleModels[1].(map[string]interface{})

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	assert.Equal(t, exampleModel1.Id, int(exampleModelResponse1["id"].(float64)))
	assert.Equal(t, exampleModel1.Name, exampleModelResponse1["name"])

	assert.Equal(t, exampleModel2.Id, int(exampleModelResponse2["id"].(float64)))
	assert.Equal(t, exampleModel2.Name, exampleModelResponse2["name"])
}

func TestGetExampleModelByIdSuccess(t *testing.T) {
	db := SetupTestDB()
	ClearExampleModelsTable(db)

	tx, _ := db.Begin()
	exampleModelRepository := repository.NewExampleModelRepository()
	exampleModel := exampleModelRepository.Save(context.Background(), tx, entity.ExampleModel{
		Name: "Something",
	})
	tx.Commit()

	router := SetupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/example-models/"+strconv.Itoa(exampleModel.Id), nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, exampleModel.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, exampleModel.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetExampleModelByIdFailed(t *testing.T) {
	db := SetupTestDB()
	ClearExampleModelsTable(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/example-models/10000", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestCreateExampleModelSuccess(t *testing.T) {
	db := SetupTestDB()
	ClearExampleModelsTable(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Something"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/example-models", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Something", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateExampleModelFailed(t *testing.T) {
	db := SetupTestDB()
	ClearExampleModelsTable(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/example-models", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
