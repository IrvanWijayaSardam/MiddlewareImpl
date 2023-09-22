package controller

import (
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/MiddlewareImpl/helper"
	"github.com/IrvanWijayaSardam/MiddlewareImpl/model"
	"github.com/labstack/echo/v4"
)

type BooksController struct {
	model model.BooksModel
}

func (bc *BooksController) InitBooksController(bm model.BooksModel) {
	bc.model = bm
}

func (bc *BooksController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var bookData model.Books
		if err := c.Bind(&bookData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request data"})
		}

		newBook := bc.model.Create(bookData)
		if newBook == nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create book"})
		}

		return c.JSON(http.StatusCreated, newBook)
	}
}

func (bc *BooksController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		bookID := id
		if bookID == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid book ID"})
		}

		book := bc.model.Get(bookID)
		if book == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
		}

		return c.JSON(http.StatusOK, book)
	}
}

func (bc *BooksController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		books := bc.model.GetAll()
		if books == nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch books"})
		}
		return c.JSON(http.StatusOK, books)
	}
}

func parseUint(s string) uint {
	id, _ := strconv.ParseUint(s, 10, 64)
	return uint(id)
}

func (bc *BooksController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")

		if paramId == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid id", nil))
		}

		var input = model.Books{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}
		input.ID = paramId

		var res = bc.model.UpdateData(input)

		if !res {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res))
	}
}

func (bc *BooksController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid id", nil))
		}

		bc.model.Delete(cnv)

		return c.JSON(http.StatusNoContent, nil)
	}
}
