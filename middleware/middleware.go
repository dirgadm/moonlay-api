package middleware

import (
	"context"
	"net/http"
	"project-version3/moonlay-api/pkg/constants"
	"project-version3/moonlay-api/pkg/ehttp"
	jwtx "project-version3/moonlay-api/pkg/jwt"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
)

// Middleware defines object for order api custom middleware
type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Authorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			authorization := ctx.Request().Header.Get("Authorization")
			var match bool
			match, err = regexp.MatchString("^Bearer .+", authorization)
			if err != nil || !match {
				return echo.ErrUnauthorized
			}

			j := jwtx.NewJWT([]byte(viper.GetString(`jwt.key`)))

			tokenStr := strings.Split(authorization, " ")

			var token *jwt.Token
			token, err = j.Parse(tokenStr[1])
			if err != nil {
				return echo.ErrUnauthorized
			}

			var claims *jwtx.UserClaim
			var ok bool
			claims, ok = token.Claims.(*jwtx.UserClaim)
			if !ok {
				return echo.ErrUnauthorized
			}

			expiresAt := claims.ExpiresAt
			if expiresAt <= time.Now().Unix() {
				ctx.JSON(http.StatusUnauthorized, ehttp.FormatResponse{
					Code:    http.StatusUnauthorized,
					Status:  "failure",
					Message: "Your token is expired",
				})
				return echo.ErrUnauthorized
			}

			ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), constants.KeyToken, token)))
			ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), constants.KeyUserID, claims.UserId)))
			return next(ctx)
		}
	}
}
