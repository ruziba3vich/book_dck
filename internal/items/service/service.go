package service

import "github.com/ruziba3vich/boock/internal/items/repository"

type (
	Service struct {
		storage repository.IBookRepo
	}
)

/*
	CreateBook(*models.CreateBookRequest) (*models.Book, error)
	UpdateBook(*models.UpdateBookRequest) (*models.Book, error)
	GetBookById(*models.GetBookByIdRequest) (*models.Book, error)
	GetAllBooks(*models.GetAllBooksRequest) (*models.GetSeveralResponse, error)
	GetBooksByAuthor(*models.GetBooksByAuthorRequest) (*models.GetSeveralResponse, error)
	GetBooksByName(*models.GetBooksByNameRequest) (*models.GetSeveralResponse, error)
	SearchBooks(*models.SearchBooksRequest) (*models.GetSeveralResponse, error)
*/
