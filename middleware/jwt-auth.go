package middleware

import (
	"net/http"
	"os"

	"github.com/IrvanWijayaSardam/MiddlewareImpl/helper"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/service"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthorizeJWT(jwtService service.JWTService) echo.MiddlewareFunc {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		// Handle error if .env file is not found or cannot be read
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				response := helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
				return c.JSON(http.StatusInternalServerError, response)
			}
		}
	}

	// Retrieve the secret key from the environment variables
	secretKey := os.Getenv("SECRET_KEY")

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(secretKey), // Use the secret key from .env
		TokenLookup: "header:XToken",   // Look for the token in the XToken header
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			c.Set("userid", claims["sub"]) // Change "userid" to match your claim key
			c.Set("issuer", claims["iss"]) // Change "issuer" to match your claim key
		},
		ErrorHandler: func(err error) error {
			if echoErr, ok := err.(*echo.HTTPError); ok {
				response := helper.BuildErrorResponse("Token is not valid", echoErr.Message.(string), nil)
				return echo.NewHTTPError(http.StatusUnauthorized, response)
			}
			return err
		},
	})
}
