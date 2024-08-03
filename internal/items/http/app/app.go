package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/boock/internal/items/http/handler"
)

func Run(router *gin.Engine, handler *handler.Handler, logger *log.Logger, host string) error {
	r := router.Group("/books")

	r.POST("/", handler.CreateBookHandler)
	r.PUT("/", handler.UpdateBookHandler)
	r.GET("/", handler.GetBookByIdHandler)
	r.GET("/", handler.GetAllBooksHandler)
	r.GET("/author", handler.GetBooksByAuthorHandler)
	r.GET("/name", handler.GetBooksByNameHandler)
	r.GET("/search", handler.SearchBooksHandler)
	r.DELETE("/", handler.DeleteBookByIdHandler)

	return router.Run(host)
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
