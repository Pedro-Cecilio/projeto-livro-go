package repositories

import "github.com/Pedro-Cecilio/projeto-livro-go/models"

type BookRepository interface {
	AddBookRead(bookRead models.BookRead)(bool, error)
	AddBookBeingRead(bookBeingRead models.BooksBeingRead) (bool, error)
	AddBookToRead(bookToRead models.BooksToRead) (bool, error)
}