package backend

import (
	"backend/application/usecase"
	"backend/infrastructure"
	"backend/infrastructure/config"
	"backend/presentation/router"
)

func main() {
    db := config.NewDB()

    tweetRepo := infrastructure.NewTweetRepository(db)
    userRepo := infrastructure.NewUserRepository(db)

    fetchUserUsecase := usecase.NewFetchUserUsecase(userRepo, tweetRepo)
    fetchUserRouter := router.NewFetchUserRouter(fetchUserUsecase)

    fetchAllTweetsUsecase := usecase.NewFetchAllTweetsUsecase(tweetRepo)
    fetchAllTweetsRouter := router.NewFetchAllTweetsRouter(fetchAllTweetsUsecase)

    createTweetUsecase := usecase.NewCreateTweetUsecase(tweetRepo)
    createTweetRouter := router.NewCreateTweetRouter(createTweetUsecase)

    router.InitRoute(fetchUserRouter, fetchAllTweetsRouter, createTweetRouter)
}

