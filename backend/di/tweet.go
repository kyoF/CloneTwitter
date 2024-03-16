package di

import (
	"backend/application/usecase"
	"backend/infrastructure/mysql"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func Tweet(db *gorm.DB) router.ITweetRouter {
	tweetRepo := mysql.NewTweetRepository(db)
	usecase := usecase.NewTweetUsecase(tweetRepo)
	return router.NewTweetRouter(usecase)
}
