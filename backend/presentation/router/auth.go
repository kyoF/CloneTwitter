package router

import (
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAuthRouter interface {
    SignUp() echo.HandlerFunc
    Login() echo.HandlerFunc
}

type authRouter struct {
    uc usecase.IAuthUsecase
}

func NewAuthRouter(uc usecase.IAuthUsecase) IAuthRouter {
    return &authRouter{uc}
}

func (r *authRouter) SignUP() echo.HandlerFunc {
    return func(ctx echo.Context) error {
		userId := ctx.Param("userId")
		email := ctx.Param("email")
		password := ctx.Param("password")
		err := r.uc.SignUp(userId, email, password)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, nil)
	}
}

func (r *authRouter) Login() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userId := ctx.Param("userId")
		email := ctx.Param("email")
		password := ctx.Param("password")
		err := r.uc.Login(email, userId, password)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, nil)
	}
}
