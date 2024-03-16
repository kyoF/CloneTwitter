package router

import (
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITweetRouter interface {
	FetchAll() echo.HandlerFunc
	Fetch() echo.HandlerFunc
	Create() echo.HandlerFunc
}

type tweetRouter struct {
	uc usecase.ITweetUsecase
}

func NewTweetRouter(uc usecase.ITweetUsecase) ITweetRouter {
	return &tweetRouter{
		uc: uc,
	}
}

func (r *tweetRouter) FetchAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tweets, err := r.uc.FetchAll()
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, tweets)
	}
}

func (r *tweetRouter) Fetch() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        tweetId := ctx.Param("tweetId")
		tweet, err := r.uc.Fetch(tweetId)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, tweet)
	}
}

func (r *tweetRouter) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
        userId := ctx.Param("userId")
        text := ctx.Param("text")
		tweet, err := r.uc.Create(userId, text)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, tweet)
	}
}
