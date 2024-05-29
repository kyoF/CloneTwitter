package di

import (
	"backend/application/query_service"
	"backend/application/usecase"
	"backend/infrastructure/mysql"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func Tweet(db *gorm.DB) router.ITweetRouter {
	userRepo := mysql.NewUserRepository(db)
	tweetRepo := mysql.NewTweetRepository(db)
	usecase := usecase.NewTweetUsecase(tweetRepo)
	queryService := query_service.NewFetchTweetCardQueryService(userRepo, tweetRepo)
	return router.NewTweetRouter(usecase, queryService)
}
