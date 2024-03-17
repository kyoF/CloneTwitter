package router

import (
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserRouter interface {
	Fetch() echo.HandlerFunc
}

type userRouter struct {
	uc usecase.IUserUsecase
}

func NewUserRouter(uc usecase.IUserUsecase) IUserRouter {
	return &userRouter{uc}
}

func (u *userRouter) Fetch() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("userId")
		user, err := u.uc.Fetch(userId)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, &user)
	}
}
