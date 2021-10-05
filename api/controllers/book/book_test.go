package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"project-api/config"
	"project-api/models"
	"project-api/util"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	// create database connection
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)

	// cleaning data before testing
	db.Migrator().DropTable(&models.Book{})
	db.AutoMigrate(&models.Book{})

	// preparate dummy data
	var newBook models.Book
	newBook.Title = "Alfabet"
	newBook.Author = "Alterra"
	newBook.Publisher = "Alterra"

	// dummy data with model
	bookModel := models.NewBookModel(db)
	_, err := bookModel.InsertBook(newBook)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetAllBookController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	bookModel := models.NewBookModel(db)
	bookController := NewController(bookModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/books")

	bookController.GetAllBookController(context)

	// build struct response
	type Response []struct {
		Title     string `json:"title"`
		Author    string `json:"author"`
		Publisher string `json:"publisher"`
	}

	var response Response
	resBody := res.Body.String()

	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /books", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, "Alfabet", response[0].Title)
		assert.Equal(t, "Alterra", response[0].Author)

		assert.Equal(t, "Alterra", response[0].Publisher)
	})
}

func TestGetBookController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	bookModel := models.NewBookModel(db)
	bookController := NewController(bookModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/books/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	bookController.GetBookController(context)

	// Unmarshal respose string to struct
	type Response struct {
		Title     string `json:"title"`
		Author    string `json:"author"`
		Publisher string `json:"publisher"`
	}

	var response Response
	resBody := res.Body.String()

	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /books/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code) // response.Data.
		assert.Equal(t, "Alfabet", response.Title)
		assert.Equal(t, "Alterra", response.Author)

		assert.Equal(t, "Alterra", response.Author)
	})
}

func TestPostBookController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	bookModel := models.NewBookModel(db)
	bookController := NewController(bookModel)

	// input controller
	reqBody, _ := json.Marshal(map[string]string{
		"title":     "Alfabet",
		"author":    "Alterra",
		"publisher": "Alterra",
	})

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/books")

	bookController.PostBookController(context)

	// build struct response
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("POST /books", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})
}

func TestEditBookController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	bookModel := models.NewBookModel(db)
	bookController := NewController(bookModel)

	// input controller
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "Alfabet",
		"email":    "Alterra",
		"password": "Alterra",
	})

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/books/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	bookController.EditBookController(context)

	// build struct response
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("PUT /books/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})
}

func TestDeleteBookController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	bookModel := models.NewBookModel(db)
	bookController := NewController(bookModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/books/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	bookController.DeleteBookController(context)

	// build struct response
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("PUT /books/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})
}
