package di

import (
	"backend/application/query_service"
	"backend/infrastructure/mysql"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func User(db *gorm.DB) router.IUserRouter {
	userRepo := mysql.NewUserRepository(db)
	tweetRepo := mysql.NewTweetRepository(db)
	userQueryService := query_service.NewUserQueryService(userRepo, tweetRepo)
	return router.NewUserRouter(userQueryService)
}
