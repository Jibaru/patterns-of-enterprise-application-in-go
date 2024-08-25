package helpers

import "github.com/jibaru/template-view/internal/domain"

// BookHelper is a helper that provides access to Book properties.
type BookHelper struct {
	book domain.Book
}

// NewBookHelper creates a new BookHelper.
func NewBookHelper(book domain.Book) *BookHelper {
	return &BookHelper{book: book}
}

// GetTitle returns the title of the book.
func (bh *BookHelper) GetTitle() string {
	return bh.book.Title
}

// GetAuthor returns the author of the book.
func (bh *BookHelper) GetAuthor() string {
	return bh.book.Author
}

// GetDescription returns the description of the book.
func (bh *BookHelper) GetDescription() string {
	return bh.book.Description
}
