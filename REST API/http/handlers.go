package http

import (
	"app/books"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	library *books.Library
}

func NewHTTPHandlers(lib *books.Library) *HTTPHandlers {
	return &HTTPHandlers{
		library: lib,
	}
}

// pattern: /library
// method: POST
// info: JSON in HTTP request body
func (h HTTPHandlers) AddBookHeader(w http.ResponseWriter, r *http.Request) {
	var bookDTO BookDTO
	err := json.NewDecoder(r.Body).Decode(&bookDTO)
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	book := books.NewBook(bookDTO.BookTitle, bookDTO.Author, bookDTO.NumberOfPages)
	h.library.AddBook(book)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		panic(err)
	}
}

// pattern: /library/{title}
// method: PATCH
// info: pattern + JSON in request body
func (h *HTTPHandlers) MarkAsReadHandler(w http.ResponseWriter, r *http.Request) {
	var readDTO ReadDTO
	err := json.NewDecoder(r.Body).Decode(&readDTO)
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if !readDTO.Complete {
		err := errors.New("complete field must be true to mark book as read")
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	pattern := mux.Vars(r)["title"]
	book, err := h.library.SetRead(pattern)
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		panic(err)
	}
}

// pattern: /library/{title}
// method: GET
// info: pattern

func (h *HTTPHandlers) GetBookHandler(w http.ResponseWriter, r *http.Request) {
	pattern := mux.Vars(r)["title"]
	book, err := h.library.GetBook(pattern)
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		panic(err)
	}
}

// pattern: /library
// method: GET
// info: Обработка GET запросов к /library с различными query параметрами
func (h *HTTPHandlers) LibraryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	
	// Проверяем наличие query параметров
	if author := query.Get("author"); author != "" {
		// Фильтрация по автору
		books := h.library.BooksByAuthor(author)
		json.NewEncoder(w).Encode(books)
		return
	}
	
	if status := query.Get("status"); status != "" {
		// Фильтрация по статусу
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			errDTO := CreateErrDTO(err)
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		}
		books := h.library.BooksByStatus(statusBool)
		json.NewEncoder(w).Encode(books)
		return
	}
	
	// Если нет query параметров, возвращаем все книги
	books := h.library.GetAllLibrary()
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		panic(err)
	}
}

// pattern: /library/{title}
// method: DELETE
// info: -
func (h *HTTPHandlers) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	err := h.library.DeleteBook(title)
	if err != nil {
		errDTO := CreateErrDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
