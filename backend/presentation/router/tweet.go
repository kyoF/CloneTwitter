package router

import (
	"backend/application/query_service"
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITweetRouter interface {
	FetchAll() echo.HandlerFunc
    FetchAllByUserId() echo.HandlerFunc
	Fetch() echo.HandlerFunc
	Create() echo.HandlerFunc
}

type tweetRouter struct {
	uc usecase.ITweetUsecase
    qs query_service.IFetchTweetCardQueryService
}

func NewTweetRouter(uc usecase.ITweetUsecase, qs query_service.IFetchTweetCardQueryService) ITweetRouter {
	return &tweetRouter{uc, qs}
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

func (r *tweetRouter) FetchAllByUserId() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userId := ctx.Param("userId")
		tweets, err := r.qs.FetchAllByUserId(userId)
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
        body := make(map[string]string)
        if err := ctx.Bind(&body); err != nil {
            ctx.Error(err)
        }
        userId := body["userId"]
        text := body["text"]
		tweet, err := r.uc.Create(userId, text)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, tweet)
	}
}
