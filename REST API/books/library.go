package books

import (
	"maps"
	"sync"
)

type Library struct {
	mtx   sync.RWMutex
	books map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}

// добавлять новые книги в нашу личную библиотеку
func (l *Library) AddBook(b Book) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	l.books[b.BookTitle] = b
}

// отмечать отдельные книги как прочитанные
func (l *Library) SetRead(bookTitle string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if book, ok := l.books[bookTitle]; !ok {
		return Book{}, ErrBookNotFound
	} else {
		book.MarkAsRead()
		l.books[bookTitle] = book
		return book, nil
	}
}

// получать информацию о какой-то конкретной книге
func (l *Library) GetBook(bookTitle string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	if book, ok := l.books[bookTitle]; ok {
		return book, nil
	}
	return Book{}, ErrBookNotFound
}

// получать список всех книг, с учётом возможной фильтрации
// по автору, прочитано/непрочитано
func (l *Library) GetAllLibrary() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	books := make(map[string]Book, len(l.books))
	maps.Copy(books, l.books)
	return books
}

func (l *Library) BooksByAuthor(author string) map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	books := make(map[string]Book)
	for k, v := range l.books {
		if v.Author == author {
			books[k] = v
		}
	}
	return books
}

func (l *Library) BooksByStatus(read bool) map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	books := make(map[string]Book)
	for k, v := range l.books {
		if v.IsRead == read {
			books[k] = v
		}
	}
	return books
}

// удалять книги из нашей библиотеки
func (l *Library) DeleteBook(bookTitle string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.books[bookTitle]; ok {
		delete(l.books, bookTitle)
		return nil
	}
	return ErrBookNotFound
}
