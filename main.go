package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"go-api-starter/app"
	"go-api-starter/helper"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	router := app.NewRouter(db, validate)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
