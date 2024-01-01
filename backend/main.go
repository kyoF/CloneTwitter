package backend

import (
	"backend/presentation/router"
	"backend/application/query_service"
	"backend/application/usecase"
	"backend/infrastructure"
	"backend/infrastructure/config"
)

func main() {
    db := config.NewDB()

    tweetRepo := infrastructure.NewTweetRepository(db)
    userRepo := infrastructure.NewUserRepository(db)

    fetchUserQueryService := query_service.NewFetchUserQueryService(userRepo, tweetRepo)
    fetchUserRouter := router.NewFetchUserRouter(fetchUserQueryService)

    fetchAllTweetsQueryService := query_service.NewFetchAllTweetsQueryService(tweetRepo)
    fetchAllTweetsRouter := router.NewFetchAllTweetsRouter(fetchAllTweetsQueryService)

    createTweetUsecase := usecase.NewCreateTweetUsecase(tweetRepo)
    createTweetRouter := router.NewCreateTweetRouter(createTweetUsecase)

    router.InitRoute(fetchUserRouter, fetchAllTweetsRouter, createTweetRouter)
}

