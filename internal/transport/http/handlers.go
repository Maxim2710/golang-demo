package http

import (
	"github.com/gin-gonic/gin"
	"golang-demo/internal/models"
	"golang-demo/internal/service"
	"net/http"
)

type BookHandler struct {
	Service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.Service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, book)
}
