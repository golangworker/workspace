// DTO = Data Transfer Object
package http

import (
	"encoding/json"
	"time"
)

// структура json запроса от пользователя при создании книги
type BookDTO struct {
	BookTitle     string    `json:"book_title"`
	Author        string    `json:"author"`
	NumberOfPages int       `json:"number_of_pages"`
	TimeAdded     time.Time `json:"time_added"`
}

func CreateBookDTO(bookTitle, author string, numberOfPages int) BookDTO {
	return BookDTO{
		BookTitle:     bookTitle,
		Author:        author,
		NumberOfPages: numberOfPages,
		TimeAdded:     time.Now(),
	}
}

// структура json ответа пользователю о произошедшей ошибке
type ErrDTO struct {
	Message string
	Time    time.Time
}

func CreateErrDTO(err error) ErrDTO {
	return ErrDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (e ErrDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// структура json запроса от пользователя для изменения статуса прочтения книги
type ReadDTO struct {
	Complete bool `json:"complete"`
}

func CreateReadDTO(b bool) ReadDTO {
	return ReadDTO{
		Complete: b,
	}
}
