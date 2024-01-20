package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"crudapp/models"
	"crudapp/services"
)

func InitializeRoutes(router *gin.Engine) {
	// Set up routes
	router.POST("/books", CreateBook)
	router.GET("/books/:id", GetBook)
	router.GET("/books", GetAllBooks)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := services.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := services.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {
	books := services.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := services.UpdateBook(id, updatedBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

