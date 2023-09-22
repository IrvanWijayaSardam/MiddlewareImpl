package controller

import (
	"net/http"

	"github.com/IrvanWijayaSardam/MiddlewareImpl/model"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/service"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	model      model.AuthModel
	jwtService service.JWTService
}

func (ac *AuthController) InitAuthController(am model.AuthModel, jwtService service.JWTService) {
	ac.model = am
	ac.jwtService = jwtService
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginData model.Login
		if err := c.Bind(&loginData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request data"})
		}

		success, err := ac.model.Login(loginData)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		}

		if !success {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		}

		// Generate a JWT token
		token := ac.jwtService.GenerateToken(loginData.Email)

		// Create a response map including the token
		responseData := map[string]interface{}{
			"message": "Login successful",
			"token":   token,
		}

		return c.JSON(http.StatusOK, responseData)

	}
}
