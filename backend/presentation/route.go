package presentation

import (
	customMiddleware "backend/presentation/middleware"
	"backend/presentation/router"
	"net/http"

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

	test := a.Group("/authenticated")
	test.Use(customMiddleware.JWTMiddleware)
	test.GET("/test", func(c echo.Context) error {
		testText := c.Get("user_id").(string)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Private data for user " + testText,
		})
	})
	a.POST("/signup", authRouter.SignUp())
	a.POST("/login", authRouter.Login())

	a.GET("/user/:userId", userRouter.Fetch())
	a.GET("/user/:userId/tweets", tweetRouter.FetchAllByUserId())

	a.GET("/tweets", tweetRouter.FetchAll())
	a.GET("/tweet/:tweetId", tweetRouter.Fetch())
	a.POST("/tweet/create", tweetRouter.Create())

	a.POST("/tweet/like", likeRouter.ToggleLike())

	e.Logger.Fatal(e.Start(":8888"))
}
