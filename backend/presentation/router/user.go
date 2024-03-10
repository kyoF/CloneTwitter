package router

import (
	"backend/application/query_service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserRouter interface {
	Fetch() echo.HandlerFunc
}

type userRouter struct {
	qs query_service.IUserQueryService
}

func NewUserRouter(qs query_service.IUserQueryService) IUserRouter {
	return &userRouter{qs}
}

func (u *userRouter) Fetch() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("userId")
		dto, err := u.qs.Fetch(userId)
		if err != nil {
			ctx.Error(err)
		}
		return ctx.JSON(http.StatusOK, dto)
	}
}
