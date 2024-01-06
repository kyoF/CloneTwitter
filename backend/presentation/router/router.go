package router

import "github.com/labstack/echo/v4"

func InitRoute(
    fetchUserRouter IFetchUserRouter,
    fetchAllTweetRouter IFetchAllTweetsRouter,
    createTweetRouter ICreateTweetRouter,
) {
    e := echo.New()
    e.GET("/user/:id", fetchUserRouter.Route())
    e.GET("/tweets", fetchAllTweetRouter.Route())
    e.POST("/user/:id/tweet/create", createTweetRouter.Route())

    e.Logger.Fatal(e.Start(":8888"))
}

