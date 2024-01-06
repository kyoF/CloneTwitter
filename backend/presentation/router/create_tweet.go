package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/application/usecase"
)

type ICreateTweetRouter interface {
    Route() echo.HandlerFunc
}

type createTweetRouter struct {
    uc usecase.ICreateTweetUsecase
}

func NewCreateTweetRouter(uc usecase.ICreateTweetUsecase) ICreateTweetRouter {
    return &createTweetRouter{uc}
}

type requestBody struct {
    UserId string `json:"userId"`
    Text string `json:"text"`
}

func (r *createTweetRouter) Route() echo.HandlerFunc {
    return func(ctx echo.Context) error {
        body := new(requestBody)
        if err := ctx.Bind(body); err != nil {
            return ctx.String(http.StatusInternalServerError, "Error!")
        }
        body.UserId = ctx.Param("id")
        tweet, err := r.uc.CreateTweet(body.UserId, body.Text)
        if err != nil {
            ctx.Error(err)
        }
        return ctx.JSON(http.StatusOK, tweet)
    }
}

