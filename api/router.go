package api

import (
	"project-api/api/controllers/book"
	"project-api/api/controllers/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userController *user.Controller) {

	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	e.POST("/users/register", userController.PostUserController)
	e.POST("/users/login", userController.DeleteUserController)

	// ------------------------------------------------------------------
	// CRUD Customer
	// ------------------------------------------------------------------
	e.GET("/users", userController.GetAllUserController)
	e.GET("/users/:id", userController.GetUserController)
	e.PUT("/users/:id", userController.EditUserController)
	e.DELETE("/users/:id", userController.DeleteUserController)

}

func RegisterPathBook(e *echo.Echo, bookController *book.Controller) {
	e.GET("/books", bookController.GetAllBookController)
	e.GET("/books/:id", bookController.GetBookController)
	e.PUT("/books/:id", bookController.EditBookController)
	e.DELETE("/books/:id", bookController.DeleteBookController)
}
