package routes

import (
	"github.com/IrvanWijayaSardam/MiddlewareImpl/controller"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/middleware"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/service"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controller.UserController) {
	e.POST("/users", uc.Create())
	e.GET("/users", uc.GetUsers())
	e.GET("/users/:id", uc.Get())
	e.DELETE("/users", uc.Delete())
	e.PUT("/users/:id", uc.Update())
}

func RouteBook(e *echo.Echo, bc controller.BooksController, jwtService service.JWTService) {
	// Use JWT authorization middleware to protect routes
	e.POST("/books", bc.Create(), middleware.AuthorizeJWT(jwtService))
	e.GET("/books", bc.GetAll(), middleware.AuthorizeJWT(jwtService))
	e.GET("/books/:id", bc.Get(), middleware.AuthorizeJWT(jwtService))
	e.DELETE("/books/:id", bc.Delete(), middleware.AuthorizeJWT(jwtService))
	e.PUT("/books/:id", bc.Update(), middleware.AuthorizeJWT(jwtService))
}

func RouteAuth(e *echo.Echo, ac controller.AuthController) {
	e.POST("login", ac.Login())
}
