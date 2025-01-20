package http

import (
	"github.com/gin-gonic/gin"
	"golang-demo/internal/service"
)

func SetupRouter(service *service.BookService) *gin.Engine {
	router := gin.New()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	bookHandler := NewBookHandler(service)

	router.POST("/books", bookHandler.CreateBook)
	router.GET("/books/:id", bookHandler.GetBookById)
	router.DELETE("/books/:id", bookHandler.DeleteBookById)

	return router
}
