package composites

import (
	"context"
	"github.com/atlant1da-404/application-architecture/pkg/client/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBComposite struct {
	db *mongo.Database
}

type MongoDBCompositeOptions struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	AuthSource string
}

func NewMongoDBComposite(ctx context.Context, opt *MongoDBCompositeOptions) (*MongoDBComposite, error) {
	client, err := mongodb.NewClient(ctx, opt.Host, opt.Port, opt.Username, opt.Password, opt.Database, opt.AuthSource)
	if err != nil {
		return nil, err
	}
	return &MongoDBComposite{client}, err
}
