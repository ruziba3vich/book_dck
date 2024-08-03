package service

import (
	"context"

	"github.com/ruziba3vich/boock/internal/items/repository"
	"github.com/ruziba3vich/boock/internal/models"
)

type (
	Service struct {
		storage repository.IBookRepo
	}
)

func New(storage repository.IBookRepo) repository.IBookRepo {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateBook(ctx context.Context, req *models.CreateBookRequest) (*models.Book, error) {
	return s.storage.CreateBook(ctx, req)
}
func (s *Service) UpdateBook(ctx context.Context, req *models.UpdateBookRequest) (*models.Book, error) {
	return s.storage.UpdateBook(ctx, req)
}
func (s *Service) GetBookById(ctx context.Context, req *models.GetBookByIdRequest) (*models.Book, error) {
	return s.storage.GetBookById(ctx, req)
}
func (s *Service) GetAllBooks(ctx context.Context, req *models.GetAllBooksRequest) (*models.GetSeveralResponse, error) {
	return s.storage.GetAllBooks(ctx, req)
}
func (s *Service) GetBooksByAuthor(ctx context.Context, req *models.GetBooksByAuthorRequest) (*models.GetSeveralResponse, error) {
	return s.storage.GetBooksByAuthor(ctx, req)
}
func (s *Service) GetBooksByName(ctx context.Context, req *models.GetBooksByNameRequest) (*models.GetSeveralResponse, error) {
	return s.storage.GetBooksByName(ctx, req)
}
func (s *Service) SearchBooks(ctx context.Context, req *models.SearchBooksRequest) (*models.GetSeveralResponse, error) {
	return s.storage.SearchBooks(ctx, req)
}
func (s *Service) DeleteBookById(ctx context.Context, req *models.DeleteBookByIdRequest) error {
	return s.storage.DeleteBookById(ctx, req)
}

/*
	CreateBook(*models.CreateBookRequest) (*models.Book, error)
	UpdateBook(*models.UpdateBookRequest) (*models.Book, error)
	GetBookById(*models.GetBookByIdRequest) (*models.Book, error)
	GetAllBooks(*models.GetAllBooksRequest) (*models.GetSeveralResponse, error)
	GetBooksByAuthor(*models.GetBooksByAuthorRequest) (*models.GetSeveralResponse, error)
	GetBooksByName(*models.GetBooksByNameRequest) (*models.GetSeveralResponse, error)
	SearchBooks(*models.SearchBooksRequest) (*models.GetSeveralResponse, error)
	DeleteBookById(context.Context,*models.DeleteBookByIdRequest) error
*/
