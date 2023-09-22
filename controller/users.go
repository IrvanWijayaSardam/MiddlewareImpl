package controller

import (
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/MiddlewareImpl/helper"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/model"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model model.UsersModel
}

func (uc *UserController) InitUserController(um model.UsersModel) {
	uc.model = um
}

func (uc *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Users{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}

		var res = uc.model.Create(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res))
	}
}

func (uc *UserController) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res = uc.model.GetAll()

		return c.JSON(http.StatusOK, helper.FormatResponse("success", res))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		if paramId == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid id", nil))
		}

		var input = model.Users{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}
		input.ID = paramId

		var res = uc.model.UpdateData(input)

		if !res {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res))
	}
}

func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid id", nil))
		}

		uc.model.Delete(cnv)

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (uc *UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		if paramId == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid id", nil))
		}
		var res = uc.model.Get(paramId)

		return c.JSON(http.StatusOK, helper.FormatResponse("success", res))
	}
}
