package di

import (
	"backend/application/query_service"
	"backend/application/usecase"
	"backend/infrastructure"
	"backend/presentation/router"

	"gorm.io/gorm"
)

func Auth(db *gorm.DB) router.IAuthRouter {
    authRepo := infrastructure.NewAuthInfra(db)
    authUsecase := usecase.NewAuthUsecase(authRepo)
    return router.NewAuthRouter(authUsecase)
}

func User(db *gorm.DB) router.IUserRouter {
    userRepo := infrastructure.NewUserRepository(db)
    tweetRepo := infrastructure.NewTweetRepository(db)
    userQueryService := query_service.NewUserQueryService(userRepo, tweetRepo)
    return router.NewUserRouter(userQueryService)
}

func Tweet(db *gorm.DB) router.ITweetRouter {
	tweetRepo := infrastructure.NewTweetRepository(db)
	queryService := query_service.NewTweetQueryService(tweetRepo)
	usecase := usecase.NewTweetUsecase(tweetRepo)
	return router.NewTweetRouter(queryService, usecase)
}
