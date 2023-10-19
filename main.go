package main

import (
	"belajar-golang-resful-api/app"
	"belajar-golang-resful-api/controller"
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/middleware"
	"belajar-golang-resful-api/repository"
	"belajar-golang-resful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMidlleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMidlleware,
	}
}

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	authMidlleware := middleware.NewAuthMiddleware(router)

	server := NewServer(authMidlleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
