package storage

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/boock/internal/items/config"
	"github.com/ruziba3vich/boock/internal/items/redisservice"
	"github.com/ruziba3vich/boock/internal/models"
)

type Storage struct {
	redis        *redisservice.RedisService
	postgres     *sql.DB
	queryBuilder sq.StatementBuilderType
	cfg          *config.Config
	logger       *log.Logger
}

func New(redis *redisservice.RedisService, postgres *sql.DB, queryBuilder sq.StatementBuilderType, cfg *config.Config, logger *log.Logger) *Storage {
	return &Storage{
		redis:        redis,
		postgres:     postgres,
		queryBuilder: queryBuilder,
		cfg:          cfg,
		logger:       logger,
	}
}

func (s *Storage) CreateBook(ctx context.Context, req *models.CreateBookRequest) (*models.Book, error) {
	bookId := uuid.New().String()
	query, args, err := s.queryBuilder.Insert(s.cfg.TableName).
		Columns(s.cfg.BookId, s.cfg.Author, s.cfg.Title, s.cfg.PublisherYear).
		Values(bookId, req.Author, req.Title, req.PublisherYear).
		ToSql()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	rows, err := s.postgres.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	book := models.Book{
		BookId:        string(bookId),
		Author:        req.Author,
		Title:         req.Title,
		PublisherYear: req.PublisherYear,
	}
	return s.redis.StoreBookInRedis(ctx, &book)
}

func (s *Storage) UpdateBook(ctx context.Context, req *models.UpdateBookRequest) (*models.Book, error) {
	queryBuilder := s.queryBuilder.Update(s.cfg.TableName)

	if len(req.Author) > 0 {
		queryBuilder = queryBuilder.Set(s.cfg.Author, req.Author)
	}
	if len(req.Title) > 0 {
		queryBuilder = queryBuilder.Set(s.cfg.Title, req.Title)
	}
	if req.PublisherYear != 0 {
		queryBuilder = queryBuilder.Set(s.cfg.PublisherYear, req.PublisherYear)
	}

	queryBuilder = queryBuilder.Where(sq.Eq{s.cfg.BookId: req.BookId})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	result, err := s.postgres.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	updatedBook, err := s.GetBookById(ctx, &models.GetBookByIdRequest{BookId: req.BookId})
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return s.redis.StoreBookInRedis(ctx, updatedBook)
}

func (s *Storage) GetBookById(ctx context.Context, req *models.GetBookByIdRequest) (*models.Book, error) {
	redisBook, _ := s.redis.GetBookFromRedis(ctx, req.BookId)
	if redisBook != nil {
		return redisBook, nil
	}
	query, args, err := s.queryBuilder.Select(s.cfg.BookId, s.cfg.Author, s.cfg.Title, s.cfg.PublisherYear).
		From(s.cfg.TableName).
		Where(sq.Eq{s.cfg.BookId: req.BookId}).
		ToSql()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	row := s.postgres.QueryRowContext(ctx, query, args...)
	var book models.Book
	if err := row.Scan(&book.BookId, &book.Author, &book.Title, &book.PublisherYear); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return &book, nil
}

func (s *Storage) GetAllBooks(ctx context.Context, req *models.GetAllBooksRequest) (*models.GetSeveralResponse, error) {
	query, args, err := s.queryBuilder.Select(s.cfg.BookId, s.cfg.Author, s.cfg.Title, s.cfg.PublisherYear).
		From(s.cfg.TableName).
		ToSql()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	rows, err := s.postgres.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookId, &book.Author, &book.Title, &book.PublisherYear); err != nil {
			s.logger.Println(err)
			return nil, err
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return &models.GetSeveralResponse{Books: books}, nil
}

func (s *Storage) GetBooksByAuthor(ctx context.Context, req *models.GetBooksByAuthorRequest) (*models.GetSeveralResponse, error) {
	query, args, err := s.queryBuilder.Select(s.cfg.BookId, s.cfg.Author, s.cfg.Title, s.cfg.PublisherYear).
		From(s.cfg.TableName).
		Where(sq.Eq{s.cfg.Author: req.Author}).
		ToSql()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	rows, err := s.postgres.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookId, &book.Author, &book.Title, &book.PublisherYear); err != nil {
			s.logger.Println(err)
			return nil, err
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return &models.GetSeveralResponse{Books: books}, nil
}

func (s *Storage) GetBooksByName(ctx context.Context, req *models.GetBooksByNameRequest) (*models.GetSeveralResponse, error) {
	query, args, err := s.queryBuilder.Select(s.cfg.BookId, s.cfg.Author, s.cfg.Title, s.cfg.PublisherYear).
		From(s.cfg.TableName).
		Where(sq.Like{s.cfg.Title: "%" + s.cfg.Title + "%"}).
		ToSql()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	rows, err := s.postgres.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookId, &book.Author, &book.Title, &book.PublisherYear); err != nil {
			s.logger.Println(err)
			return nil, err
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return &models.GetSeveralResponse{Books: books}, nil
}

func (s *Storage) SearchBooks(*models.SearchBooksRequest) (*models.GetSeveralResponse, error) {
	return nil, nil
}

/*
	CreateBook(*models.CreateBookRequest) (*models.Book, error) ---
	UpdateBook(*models.UpdateBookRequest) (*models.Book, error)
	GetBookById(*models.GetBookByIdRequest) (*models.Book, error)
	GetAllBooks(*models.GetAllBooksRequest) (*models.GetSeveralResponse, error)
	GetBooksByAuthor(*models.GetBooksByAuthorRequest) (*models.GetSeveralResponse, error)
	GetBooksByName(*models.GetBooksByNameRequest) (*models.GetSeveralResponse, error)
	SearchBooks(*models.SearchBooksRequest) (*models.GetSeveralResponse, error)
*/
