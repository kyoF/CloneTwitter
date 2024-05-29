package di

import (
	"backend/application/usecase"
	"backend/infrastructure/mysql"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func Like(db *gorm.DB) router.ILikeRouter {
	likeRepo := mysql.NewLikeInfra(db)
	tweetRepo := mysql.NewTweetRepository(db)
	likeUsecase := usecase.NewLikeUsecase(likeRepo, tweetRepo)
	return router.NewLikeRouter(likeUsecase)
}
