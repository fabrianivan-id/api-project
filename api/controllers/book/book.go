package book

import (
	"net/http"
	"strconv"

	"project-api/api/common"
	"project-api/models"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	bookModel models.BookModel
}

func NewController(bookModel models.BookModel) *Controller {
	return &Controller{
		bookModel,
	}
}

func (controller *Controller) GetAllBookController(c echo.Context) error {
	book, err := controller.bookModel.GetAllBook()
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, book)
}

func (controller *Controller) GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	book, err := controller.bookModel.GetBook(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	response := GetBookResponse{
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) PostBookController(c echo.Context) error {
	// bind request value
	var bookRequest PostBookRequest

	if err := c.Bind(&bookRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	book := models.Book{
		Title:     bookRequest.Title,
		Author:    bookRequest.Author,
		Publisher: bookRequest.Publisher,
	}

	_, err := controller.bookModel.InsertBook(book)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) EditBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	// bind request value
	var bookRequest EditBookRequest
	if err := c.Bind(&bookRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	book := models.Book{
		Title:     bookRequest.Title,
		Author:    bookRequest.Author,
		Publisher: bookRequest.Publisher,
	}

	if _, err := controller.bookModel.EditBook(book, id); err != nil {
		return c.JSON(http.StatusNotFound, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	if _, err := controller.bookModel.DeleteBook(id); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}
