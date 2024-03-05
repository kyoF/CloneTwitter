package router

import "github.com/labstack/echo/v4"

func InitRoute(
    signUpRouter ISignUpRouter,
    loginRouter ILoginRouter,
    fetchUserRouter IFetchUserRouter,
    fetchAllTweetRouter IFetchAllTweetsRouter,
    createTweetRouter ICreateTweetRouter,
) {
    e := echo.New()
    e.GET("/signup", signUpRouter.Route())
    e.GET("/login", loginRouter.Route())
    e.GET("/user/:id", fetchUserRouter.Route())
    e.GET("/tweets", fetchAllTweetRouter.Route())
    e.POST("/user/:id/tweet/create", createTweetRouter.Route())

    e.Logger.Fatal(e.Start(":8888"))
}

