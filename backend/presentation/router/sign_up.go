package router

import (
	"backend/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ISignUpRouter interface {
	Route() echo.HandlerFunc
}

type signUpRouter struct {
	uc usecase.ISignUpUsecase
}

func NewSignUpRouter(uc usecase.ISignUpUsecase) ISignUpRouter {
	return &signUpRouter{uc}
}

func (r *signUpRouter) Route() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("userId")
		name := ctx.Param("name")
		email := ctx.Param("email")
		password := ctx.Param("password")
		err := r.uc.SignUp(userId, name, email, password)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, nil)
	}
}
