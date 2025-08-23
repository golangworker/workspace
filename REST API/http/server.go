package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	handlers *HTTPHandlers
}

func NewHTTPServer(h *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: h,
	}
}

func (s *HTTPServer) Start() error {
	router := mux.NewRouter()
	router.Path("/library").Methods("POST").HandlerFunc(s.handlers.AddBookHeader)
	router.Path("/library/{title}").Methods("PATCH").HandlerFunc(s.handlers.MarkAsReadHandler)
	router.Path("/library/{title}").Methods("GET").HandlerFunc(s.handlers.GetBookHandler)
	router.Path("/library").Methods("GET").HandlerFunc(s.handlers.LibraryHandler)
	router.Path("/library/{title}").Methods("DELETE").HandlerFunc(s.handlers.DeleteHandler)

	log.Println("Server started on port 8080")
	return http.ListenAndServe("localhost:8080", router)
}
