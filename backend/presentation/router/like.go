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
        body := make(map[string]string)
        if err := ctx.Bind(&body); err != nil {
            ctx.Error(err)
        }
		userId := body["userId"]
		tweetId := body["tweetId"]
		likesCount, err := r.uc.ToggleLike(userId, tweetId)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, likesCount)
	}
}
