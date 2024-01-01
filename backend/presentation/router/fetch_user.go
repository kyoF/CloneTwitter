package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/application/usecase"
	"backend/domain/entity"
)

type IFetchUserRouter interface {
    Route() echo.HandlerFunc
}

type fetchUserRouter struct {
    uc usecase.IFetchUserUsecase
}

func NewFetchUserRouter(uc usecase.IFetchUserUsecase) IFetchUserRouter {
    return &fetchUserRouter{uc}
}

func (r *fetchUserRouter) Route() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userId := ctx.Param("id")
        user, tweet, err := r.uc.FetchUser(userId)
        if err != nil {
            ctx.Error(err)
        }
        type resType struct {
            user *entity.User
            tweets []entity.Tweet
        }
        res := resType{user: user, tweets: tweet}
        return ctx.JSON(http.StatusOK, res)
    }
}

