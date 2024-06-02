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

func (r *authRouter) SignUp() echo.HandlerFunc {
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
	type req struct {
		UserId   string `json:"userId"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(ctx echo.Context) error {
		var req req
		if err := ctx.Bind(&req); err != nil {
			ctx.Error(err)
		}
		err := r.uc.Login(req.Email, req.UserId, req.Password)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, nil)
	}
}
