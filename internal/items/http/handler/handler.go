package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/boock/internal/items/repository"
	"github.com/ruziba3vich/boock/internal/models"
)

type (
	Handler struct {
		service repository.IBookRepo
		logger  *log.Logger
	}
)

func New(service repository.IBookRepo, logger *log.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateBookHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN CreateBookHandler --")

	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.CreateBook(context.Background(), &req)
	if err != nil {
		h.logger.Println("Error creating book:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, book)
}

func (h *Handler) UpdateBookHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN UpdateBookHandler --")

	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.BookId = c.Param("id")

	book, err := h.service.UpdateBook(context.Background(), &req)
	if err != nil {
		h.logger.Println("Error updating book:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func (h *Handler) GetBookByIdHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN GetBookByIdHandler --")

	bookId := c.Param("id")

	req := &models.GetBookByIdRequest{
		BookId: bookId,
	}
	book, err := h.service.GetBookById(context.Background(), req)
	if err != nil {
		h.logger.Println("Error getting book by ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func (h *Handler) GetAllBooksHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN GetAllBooksHandler --")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.logger.Println("Error converting page to int:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.logger.Println("Error converting limit to int:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	req := &models.GetAllBooksRequest{
		Page: page, Limit: limit,
	}
	response, err := h.service.GetAllBooks(context.Background(), req)
	if err != nil {
		h.logger.Println("Error getting all books:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) GetBooksByAuthorHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN GetBooksByAuthorHandler --")

	author := c.Query("author")
	if author == "" {
		h.logger.Println("Author query parameter is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author query parameter is required"})
		return
	}

	req := &models.GetBooksByAuthorRequest{
		Author: author,
	}
	response, err := h.service.GetBooksByAuthor(context.Background(), req)
	if err != nil {
		h.logger.Println("Error getting books by author:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) GetBooksByNameHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN GetBooksByNameHandler --")

	name := c.Query("name")
	if name == "" {
		h.logger.Println("Name query parameter is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name query parameter is required"})
		return
	}

	req := &models.GetBooksByNameRequest{
		BookName: name,
	}
	response, err := h.service.GetBooksByName(context.Background(), req)
	if err != nil {
		h.logger.Println("Error getting books by name:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) SearchBooksHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN SearchBooksHandler --")

	var req models.SearchBooksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.SearchBooks(context.Background(), &req)
	if err != nil {
		h.logger.Println("Error searching books:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) DeleteBookByIdHandler(c *gin.Context) {
	h.logger.Println("-- RECEIVED A REQUEST IN DeleteBookByIdHandler --")

	bookId := c.Param("id")

	req := &models.DeleteBookByIdRequest{
		BookId: bookId,
	}
	err := h.service.DeleteBookById(context.Background(), req)
	if err != nil {
		h.logger.Println("Error deleting book by ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

/*
	CreateBook(context.Context, *models.CreateBookRequest) (*models.Book, error)
	UpdateBook(context.Context, *models.UpdateBookRequest) (*models.Book, error)
	GetBookById(context.Context, *models.GetBookByIdRequest) (*models.Book, error)
	GetAllBooks(context.Context, *models.GetAllBooksRequest) (*models.GetSeveralResponse, error)
	GetBooksByAuthor(context.Context, *models.GetBooksByAuthorRequest) (*models.GetSeveralResponse, error)
	GetBooksByName(context.Context, *models.GetBooksByNameRequest) (*models.GetSeveralResponse, error)
	SearchBooks(context.Context, *models.SearchBooksRequest) (*models.GetSeveralResponse, error)
	DeleteBookById(context.Context,*models.DeleteBookByIdRequest) error
*/
