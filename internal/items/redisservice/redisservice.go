package redisservice

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/boock/internal/models"
)

type (
	RedisService struct {
		redisDb *redis.Client
		logger  *log.Logger
	}
)

func New(redisDb *redis.Client, logger *log.Logger) *RedisService {
	return &RedisService{
		logger:  logger,
		redisDb: redisDb,
	}
}

func (r *RedisService) StoreBookInRedis(ctx context.Context, book *models.Book) (*models.Book, error) {
	byteData, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	if err := r.redisDb.Set(ctx, book.BookId, byteData, time.Hour*24).Err(); err != nil {
		return nil, err
	}
	return book, nil
}

func (r *RedisService) GetBookFromRedis(ctx context.Context, bookId string) (*models.Book, error) {
	redisBook, err := r.redisDb.Get(ctx, bookId).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		r.logger.Printf("ERROR WHILE GETTING DATA FROM REDIS : %s\n", err.Error())
		return nil, err
	}

	var book models.Book
	err = json.Unmarshal([]byte(redisBook), &book)
	if err != nil {
		r.logger.Printf("ERROR WHILE MARSHALING DATA : %s\n", err.Error())
		return nil, err
	}
	return &book, nil
}

func (r *RedisService) DeleteBookFromRedis(ctx context.Context, bookId string) error {
	result, err := r.redisDb.Del(ctx, bookId).Result()
	if err != nil {
		return err
	}

	if result == 0 {
		r.logger.Printf("User with ID %s does not exist in Redis", bookId)
	} else {
		r.logger.Printf("User with ID %s has been deleted from Redis", bookId)
	}

	return nil
}
