package services

import "crudapp/models"

func CreateBook(book models.Book) (models.Book, error) {
	return repositories.CreateBook(book)
}

func GetBook(id string) (models.Book, error) {
	return repositories.GetBook(id)
}

func GetAllBooks() []models.Book {
	return repositories.GetAllBooks()
}

func UpdateBook(id string, updatedBook models.Book) (models.Book, error) {
	return repositories.UpdateBook(id, updatedBook)
}

func DeleteBook(id string) error {
	return repositories.DeleteBook(id)
}

