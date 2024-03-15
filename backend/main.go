package main

import (
	"backend/di"
	"backend/infrastructure/config"
	"backend/presentation"
)

func main() {
	db := config.NewDB()

	presentation.InitRoute(
		di.Auth(db),
		di.User(db),
		di.Tweet(db),
		di.Like(db),
	)
}
