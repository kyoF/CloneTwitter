package di

import (
	"backend/application/query_service"
	"backend/infrastructure"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func User(db *gorm.DB) router.IUserRouter {
	userRepo := infrastructure.NewUserRepository(db)
	tweetRepo := infrastructure.NewTweetRepository(db)
	userQueryService := query_service.NewUserQueryService(userRepo, tweetRepo)
	return router.NewUserRouter(userQueryService)
}
