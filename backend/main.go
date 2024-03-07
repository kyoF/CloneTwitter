package main

import (
	"backend/application/query_service"
	"backend/application/usecase"
	"backend/infrastructure"
	"backend/infrastructure/config"
	"backend/presentation/router"
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

	signUpUsecase := usecase.NewSignUpUsecase(userRepo)
	signUpRouter := router.NewSignUpRouter(signUpUsecase)

	loginQueryService := query_service.NewLoginQueryService(userRepo)
	loginRouter := router.NewLoginRouter(loginQueryService)

	router.InitRoute(
		fetchUserRouter,
		fetchAllTweetsRouter,
		createTweetRouter,
		loginRouter,
		signUpRouter,
	)
}
