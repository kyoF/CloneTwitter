package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/application/usecase"
)

type IFetchAllTweetsRouter interface {
    Route() echo.HandlerFunc
}

type fetchAllTweetsRouter struct {
    uc usecase.IFetchAllTweetsUsecase
}

func NewFetchAllTweetsRouter(uc usecase.IFetchAllTweetsUsecase) IFetchAllTweetsRouter {
    return &fetchAllTweetsRouter{uc}
}

func (r *fetchAllTweetsRouter) Route() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userId := ctx.Param("id")
        tweets, err := r.uc.FetchAllTweets(userId)
        if err != nil {
            ctx.Error(err)
        }
        return ctx.JSON(http.StatusOK, tweets)
    }
}

