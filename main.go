package main

import (
	"golang_resfull_api/app"
	"golang_resfull_api/controller"
	"golang_resfull_api/helper"
	"golang_resfull_api/middleware"
	"golang_resfull_api/repository"
	"golang_resfull_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New() 
	categoryRepository := repository.NewCategoryRepository() 
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService) 
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}