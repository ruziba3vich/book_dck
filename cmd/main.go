package main

import (
	"log"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/boock/internal/items/config"
	"github.com/ruziba3vich/boock/internal/items/http/app"
	"github.com/ruziba3vich/boock/internal/items/http/handler"
	"github.com/ruziba3vich/boock/internal/items/redisservice"
	"github.com/ruziba3vich/boock/internal/items/service"
	"github.com/ruziba3vich/boock/internal/items/storage"
	redisCl "github.com/ruziba3vich/boock/internal/pkg/redis"
)

func main() {
	config, err := config.New()
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		logger.Fatalln(err)
	}

	db, err := storage.ConnectDB(config)
	if err != nil {
		logger.Fatalln(err)
	}

	redis, err := redisCl.NewRedisDB(config)
	if err != nil {
		logger.Fatalln(err)
	}

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	handler := handler.New(
		service.New(
			storage.New(
				redisservice.New(
					redis,
					logger,
				),
				db,
				sqrl,
				config,
				logger,
			),
		), logger,
	)

	logger.Fatalln(app.Run(gin.Default(), handler, logger, config.Server.Port))
}
