package app

import (
	"context"
	"github.com/atlant1da-404/application-architecture/internal/composites"
	"github.com/atlant1da-404/application-architecture/internal/config"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	ctx := context.Background()
	cfg := config.Get()
	router := gin.New()

	mongoComposite, err := composites.NewMongoDBComposite(ctx, &composites.MongoDBCompositeOptions{
		Host: cfg.Mongo.Host, Port: cfg.Mongo.Port, Username: cfg.Mongo.Username,
		Password: cfg.Mongo.Password, Database: cfg.Mongo.Database, AuthSource: cfg.Mongo.AuthSource,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	bookComposite, err := composites.NewBookComposite(mongoComposite)
	if err != nil {
		log.Fatalln(err.Error())
	}
	bookComposite.Handler.Register(router)

	if err := router.Run(cfg.App.BaseURL); err != nil {
		log.Fatalln(err.Error())
	}
}
