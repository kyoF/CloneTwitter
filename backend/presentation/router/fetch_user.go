package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/application/query_service"
)

type IFetchUserRouter interface {
    Route() echo.HandlerFunc
}

type fetchUserRouter struct {
    qs query_service.IFetchUserQueryService
}

func NewFetchUserRouter(qs query_service.IFetchUserQueryService) IFetchUserRouter {
    return &fetchUserRouter{qs}
}

func (r *fetchUserRouter) Route() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        userId := ctx.Param("id")
        dto, err := r.qs.FetchUser(userId)
        if err != nil {
            ctx.Error(err)
        }
        return ctx.JSON(http.StatusOK, dto)
    }
}

