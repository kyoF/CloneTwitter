package router

import (
	"backend/application/query_service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ILoginRouter interface {
	Route() echo.HandlerFunc
}

type loginRouter struct {
	qs query_service.ILoginQueryService
}

func NewLoginRouter(qs query_service.ILoginQueryService) ILoginRouter {
	return &loginRouter{qs}
}

func (r *loginRouter) Route() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		email := ctx.Param("email")
		password := ctx.Param("password")
		err := r.qs.Login(email, password)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, nil)
	}
}
