package di

import (
	"backend/application/query_service"
	"backend/application/usecase"
	"backend/infrastructure/mysql"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func Tweet(db *gorm.DB) router.ITweetRouter {
	tweetRepo := mysql.NewTweetRepository(db)
	queryService := query_service.NewTweetQueryService(tweetRepo)
	usecase := usecase.NewTweetUsecase(tweetRepo)
	return router.NewTweetRouter(queryService, usecase)
}
