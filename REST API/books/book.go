package books

import (
	"time"
)

type Book struct {
	BookTitle     string     `json:"book_title"`
	Author        string     `json:"author"`
	NumberOfPages int        `json:"number_of_pages"`
	IsRead        bool       `json:"is_read"`
	TimeAdded     time.Time  `json:"time_added"`
	TimeRead      *time.Time `json:"time_read"`
}

func NewBook(bookTitle, author string, numberOfPages int) Book {
	return Book{
		BookTitle:     bookTitle,
		Author:        author,
		NumberOfPages: numberOfPages,
		IsRead:        false,
		TimeAdded:     time.Now(),
		TimeRead:      nil,
	}
}

func (b *Book) MarkAsRead() {
	markAsReadTime := time.Now()
	b.IsRead = true
	b.TimeRead = &markAsReadTime
}
