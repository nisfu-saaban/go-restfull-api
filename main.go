package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nisfu-saaban/go-restfull-api/app"
	"github.com/nisfu-saaban/go-restfull-api/controler"
	"github.com/nisfu-saaban/go-restfull-api/helper"
	"github.com/nisfu-saaban/go-restfull-api/middleware"
	"github.com/nisfu-saaban/go-restfull-api/repository"
	"github.com/nisfu-saaban/go-restfull-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controler.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
