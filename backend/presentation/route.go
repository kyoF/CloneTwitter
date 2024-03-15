package presentation

import (
	"backend/presentation/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(
	authRouter router.IAuthRouter,
	userRouter router.IUserRouter,
	tweetRouter router.ITweetRouter,
    likeRouter router.ILikeRouter,
) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    a := e.Group("/api")

    a.POST("/signup", authRouter.SignUp())
	a.POST("/login", authRouter.Login())

    a.GET("/user/:userId", userRouter.Fetch())

    a.GET("/tweets", tweetRouter.FetchAll())
    a.GET("/tweet/:tweetId", tweetRouter.Fetch())
    a.POST("/tweet/create", tweetRouter.Create())

    a.POST("/tweet/like", likeRouter.ToggleLike())

	e.Logger.Fatal(e.Start(":8888"))
}
