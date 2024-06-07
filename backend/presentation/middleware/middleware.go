package middleware

import (
	"backend/utils"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var jwtSecretKey = []byte(utils.GetEnv("JWT_SECRET_KEY", "jwtsecretkey"))

// JWTミドルウェア
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストヘッダからAuthorizationヘッダを取得
		authHeader := c.Request().Header.Get("Authorization")

		// Bearerトークンを抽出
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return echo.ErrUnauthorized
		}
		tokenString := tokenParts[1]

		// JWTトークンを解析
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// トークンの署名アルゴリズムをチェック
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			// シークレットキーを返す
			return jwtSecretKey, nil
		})
		if err != nil {
			return echo.ErrUnauthorized
		}

		// トークンが有効かチェック
		if !token.Valid {
			return echo.ErrUnauthorized
		}

		// トークンのクレームからデータを取り出す
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		// クレームからユーザーIDを取得
		userID, ok := claims["user_id"].(string)
		if !ok {
			return echo.ErrUnauthorized
		}

		// リクエストコンテキストにユーザーIDを格納
		c.Set("user_id", userID)

		// 次のハンドラを呼び出す
		return next(c)
	}
}
