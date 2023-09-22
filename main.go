package main

import (
	"fmt"

	"github.com/IrvanWijayaSardam/MiddlewareImpl/configs"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/controller"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/model"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/routes"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()
	jwtService := service.NewJWTService()

	db := model.InitModel(*config)
	model.Migrate(db)

	bookModel := model.BooksModel{}
	bookModel.Init(db)

	bookControll := controller.BooksController{}
	bookControll.InitBooksController(bookModel)

	authModel := model.AuthModel{}
	authModel.Init(db)

	authControll := controller.AuthController{}
	authControll.InitAuthController(authModel, jwtService)

	userModel := model.UsersModel{}
	userModel.Init(db)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userControll)
	routes.RouteAuth(e, authControll)
	routes.RouteBook(e, bookControll, jwtService)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
