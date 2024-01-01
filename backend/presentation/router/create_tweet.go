package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/application/usecase"
)

type ICreateTweetRouter interface {
    Route() echo.HandlerFunc
}

type createTweetRouter struct {
    uc usecase.ICreateTweetUsecase
}

func NewCreateTweetRouter(uc usecase.ICreateTweetUsecase) ICreateTweetRouter {
    return &createTweetRouter{uc}
}

func (r *createTweetRouter) Route() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userId := ctx.Param("id")
        text := ctx.FormValue("text")
        tweet, err := r.uc.CreateTweet(userId, text)
        if err != nil {
            ctx.Error(err)
        }
        return ctx.JSON(http.StatusOK, tweet)
    }
}

