package composites

import (
	"github.com/atlant1da-404/application-architecture/internal/adapters/database/mongodb"
	"github.com/atlant1da-404/application-architecture/internal/adapters/transport/api"
	"github.com/atlant1da-404/application-architecture/internal/adapters/transport/api/v1/book"
	"github.com/atlant1da-404/application-architecture/internal/domain/service"
)

type BookComposite struct {
	Storage service.BookStorage
	Service service.BookService
	Handler api.Handler
}

func NewBookComposite(composite *MongoDBComposite) (*BookComposite, error) {
	storage := mongodb.NewBookStorage(composite.db)
	srv := service.NewBookService(storage)
	handler := book.NewHandler(srv)
	return &BookComposite{
		Storage: storage,
		Service: srv,
		Handler: handler,
	}, nil
}
