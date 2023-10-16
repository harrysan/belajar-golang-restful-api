package main

import (
	"belajar-golang-resful-api/app"
	"belajar-golang-resful-api/controller"
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/middleware"
	"belajar-golang-resful-api/repository"
	"belajar-golang-resful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
