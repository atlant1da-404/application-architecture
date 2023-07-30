package service

import (
	"context"
	"fmt"
	"github.com/atlant1da-404/application-architecture/internal/domain/entity"
)

type bookService struct {
	bookStorage BookStorage
}

func NewBookService(storage BookStorage) BookService {
	return &bookService{storage}
}

func (b *bookService) GetOne(ctx context.Context, id string) (*entity.Book, error) {
	book, err := b.bookStorage.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, fmt.Errorf("book not found")
	}

	return book, nil
}

func (b *bookService) GetAll(ctx context.Context) ([]entity.Book, error) {
	books, err := b.bookStorage.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if books == nil {
		return nil, fmt.Errorf("books not found")
	}

	return books, nil
}
