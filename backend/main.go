package main

import (
	"backend/di"
	"backend/infrastructure/config"
	"backend/presentation/router"
)

func main() {
	db := config.NewDB()

	router.InitRoute(
        di.Auth(db),
        di.User(db),
        di.Tweet(db),
	)
}
