package service

import (
	"context"
	"github.com/atlant1da-404/application-architecture/internal/domain/entity"
)

type BookService interface {
	GetOne(ctx context.Context, id string) (*entity.Book, error)
	GetAll(ctx context.Context) ([]entity.Book, error)
}
