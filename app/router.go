package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nisfu-saaban/go-restfull-api/controler"
	"github.com/nisfu-saaban/go-restfull-api/exception"
)

func NewRouter(categoryController controler.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
