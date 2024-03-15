package router

import (
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ILikeRouter interface {
	ToggleLike() echo.HandlerFunc
}

type likeRouter struct {
	uc usecase.ILikeUsecase
}

func NewLikeRouter(uc usecase.ILikeUsecase) ILikeRouter {
	return &likeRouter{uc}
}

func (r *likeRouter) ToggleLike() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("userId")
		tweetId := ctx.Param("tweetId")
		err := r.uc.ToggleLike(userId, tweetId)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, nil)
	}
}
