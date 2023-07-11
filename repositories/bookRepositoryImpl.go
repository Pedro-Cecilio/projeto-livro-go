package repositories

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/database"
	"github.com/Pedro-Cecilio/projeto-livro-go/models"
)

type BookRepositoryImpl struct{}

func (b *BookRepositoryImpl) AddBookRead(bookRead models.BookRead) (bool, error) {
	database.CreateConnection()
	result := database.DB.Create(&bookRead)
	if result.RowsAffected == 0 {
		return false, result.Error
	}
	return true, nil
}
func (b *BookRepositoryImpl) AddBookBeingRead(bookBeingRead models.BooksBeingRead) (bool, error) {
	database.CreateConnection()
	result := database.DB.Create(&bookBeingRead)
	if result.RowsAffected == 0 {
		return false, result.Error
	}
	return true, nil
}
func (b *BookRepositoryImpl) AddBookToRead(bookToRead models.BooksToRead) (bool, error) {
	database.CreateConnection()
	result := database.DB.Create(&bookToRead)
	if result.RowsAffected == 0 {
		return false, result.Error
	}
	return true, nil
}
