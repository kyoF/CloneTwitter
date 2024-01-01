package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/application/query_service"
)

type IFetchAllTweetsRouter interface {
    Route() echo.HandlerFunc
}

type fetchAllTweetsRouter struct {
    qs query_service.IFetchAllTweetsQueryService
}

func NewFetchAllTweetsRouter(qs query_service.IFetchAllTweetsQueryService) IFetchAllTweetsRouter {
    return &fetchAllTweetsRouter{qs}
}

func (r *fetchAllTweetsRouter) Route() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userId := ctx.Param("id")
        tweets, err := r.qs.FetchAllTweets(userId)
        if err != nil {
            ctx.Error(err)
        }
        return ctx.JSON(http.StatusOK, tweets)
    }
}

