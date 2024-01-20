package repositories

import (
	"errors"
	"sync"
	"crudapp/models"
)

var (
	booksDB  = make(map[string]models.Book)
	booksDBMutex sync.Mutex
)

func CreateBook(book models.Book) (models.Book, error) {
	booksDBMutex.Lock()
	defer booksDBMutex.Unlock()

	book.ID = generateID()
	booksDB[book.ID] = book

	return book, nil
}

func GetBook(id string) (models.Book, error) {
	booksDBMutex.Lock()
	defer booksDBMutex.Unlock()

	book, exists := booksDB[id]
	if !exists {
		return models.Book{}, errors.New("book not found")
	}

	return book, nil
}

func GetAllBooks() []models.Book {
	booksDBMutex.Lock()
	defer booksDBMutex.Unlock()

	books := make([]models.Book, 0, len(booksDB))
	for _, book := range booksDB {
		books = append(books, book)
	}

	return books
}

func UpdateBook(id string, updatedBook models.Book) (models.Book, error) {
	booksDBMutex.Lock()
	defer booksDBMutex.Unlock()

	_, exists := booksDB[id]
	if !exists {
		return models.Book{}, errors.New("book not found")
	}

	updatedBook.ID = id
	booksDB[id] = updatedBook

	return updatedBook, nil
}

func DeleteBook(id string) error {
	booksDBMutex.Lock()
	defer booksDBMutex.Unlock()

	_, exists := booksDB[id]
	if !exists {
		return errors.New("book not found")
	}

	delete(booksDB, id)
	return nil
}

func generateID() string {
	// Implement your own ID generation logic (e.g., UUID)
	return "some_generated_id"
}

