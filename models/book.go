package models

import "gorm.io/gorm"

// Model Customer

type Book struct {
	gorm.Model
	Title  string
	Author string
	//Gender   string `sql:"type:ENUM('male', 'female')"`
	Publisher string

	Token string `gorm:"<-:false"`
}

type GormBookModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) *GormBookModel {
	return &GormBookModel{db: db}
}

// Interface Customer

type BookModel interface {
	GetAllBook() ([]Book, error)
	GetBook(bookId int) (Book, error)
	InsertBook(Book) (Book, error)
	EditBook(book Book, bookId int) (Book, error)
	DeleteBook(bookId int) (Book, error)
}

func (m *GormBookModel) GetAllBook() ([]Book, error) {
	var book []Book
	if err := m.db.Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (m *GormBookModel) GetBook(bookId int) (Book, error) {
	var book Book
	if err := m.db.Find(&book, bookId).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) InsertBook(book Book) (Book, error) {
	if err := m.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) EditBook(newBook Book, bookId int) (Book, error) {
	var book Book
	if err := m.db.Find(&book, "id=?", bookId).Error; err != nil {
		return book, err
	}

	book.Title = newBook.Title
	book.Author = newBook.Author
	book.Publisher = newBook.Publisher

	if err := m.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (m *GormBookModel) DeleteBook(bookId int) (Book, error) {
	var book Book
	if err := m.db.Find(&book, "id=?", bookId).Error; err != nil {
		return book, err
	}
	if err := m.db.Delete(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
