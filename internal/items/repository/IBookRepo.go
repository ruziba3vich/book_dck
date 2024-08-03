package repository

import (
	"context"

	"github.com/ruziba3vich/boock/internal/models"
)

type (
	IBookRepo interface {
		CreateBook(context.Context, *models.CreateBookRequest) (*models.Book, error)
		UpdateBook(context.Context, *models.UpdateBookRequest) (*models.Book, error)
		GetBookById(context.Context, *models.GetBookByIdRequest) (*models.Book, error)
		GetAllBooks(context.Context, *models.GetAllBooksRequest) (*models.GetSeveralResponse, error)
		GetBooksByAuthor(context.Context, *models.GetBooksByAuthorRequest) (*models.GetSeveralResponse, error)
		GetBooksByName(context.Context, *models.GetBooksByNameRequest) (*models.GetSeveralResponse, error)
		SearchBooks(context.Context, *models.SearchBooksRequest) (*models.GetSeveralResponse, error)
		DeleteBookById(context.Context,*models.DeleteBookByIdRequest) error
	}
)
