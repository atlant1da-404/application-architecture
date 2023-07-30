package mongodb

import (
	"context"
	"github.com/atlant1da-404/application-architecture/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookStorage struct {
	db *mongo.Database
}

func NewBookStorage(db *mongo.Database) *bookStorage {
	return &bookStorage{db: db}
}

func (bs *bookStorage) GetOne(сtx context.Context, id string) (*entity.Book, error) {
	return nil, nil
}
func (bs *bookStorage) GetAll(ctx context.Context) ([]entity.Book, error) {
	return nil, nil
}
