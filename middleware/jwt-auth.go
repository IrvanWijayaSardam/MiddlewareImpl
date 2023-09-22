package middleware

import (
	"net/http"
	"strings"

	"github.com/IrvanWijayaSardam/MiddlewareImpl/helper"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthorizeJWT(jwtService service.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("XToken")
			if authHeader == "" {
				response := helper.BuildErrorResponse("Failed to process request", "No Token Found!", nil)
				return c.JSON(http.StatusBadRequest, response)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				c.Set("userid", claims["userid"])
				c.Set("issuer", claims["issuer"])
				return next(c)
			}

			return c.NoContent(http.StatusUnauthorized)
		}
	}
}
