package book

import (
	"github.com/atlant1da-404/application-architecture/internal/adapters/transport/api"
	"github.com/atlant1da-404/application-architecture/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	bookService service.BookService
}

func NewHandler(bookService service.BookService) api.Handler {
	return &handler{bookService: bookService}
}

const (
	getAllBooksURL = "/books"
	getBookURL     = "/book/:book_id"
)

func (h *handler) Register(router *gin.Engine) {
	router.GET(getAllBooksURL, wrapHandler(h.getAllBooks))
	router.GET(getBookURL, wrapHandler(h.getBook))
}

func (h *handler) getAllBooks(c *gin.Context) (interface{}, *httpResponseError) {
	books, err := h.bookService.GetAll(c)
	if err != nil {
		return nil, &httpResponseError{Type: ErrorTypeServer, Message: "failed to list books", Details: err}
	}

	return books, nil
}

func (h *handler) getBook(c *gin.Context) (interface{}, *httpResponseError) {
	bookId := c.Param("book_id")
	if bookId == "" {
		return nil, &httpResponseError{Type: ErrorTypeServer, Message: "book_id not found"}
	}

	book, err := h.bookService.GetOne(c, bookId)
	if err != nil {
		return nil, &httpResponseError{Type: ErrorTypeServer, Message: "failed to get books", Details: err}
	}

	return book, nil
}
