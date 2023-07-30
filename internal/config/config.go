package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type (
	// Config represents configuration for the application.
	Config struct {
		App   App
		Mongo MongoDB
	}

	// App represents general app configuration.
	App struct {
		BaseURL string `env:"APP_BASE_URL"    env-default:"localhost:8080"`
		URLLen  int    `env:"URL_LEN"    env-default:"5"`
	}

	// MongoDB represents mongoDb configuration.
	MongoDB struct {
		Host       string `env:"DB_HOST" env-default:"localhost"`
		Port       string `env:"DB_PORT" env-default:"6379"`
		Username   string `env:"DB_PORT" env-default:"db"`
		Database   string `env:"DB_DATABASE" env-default:"db"`
		Password   string `env:"DB_PASSWORD" env-default:""`
		AuthSource string `env:"DB_AUTH_SOURCE" env-default:""`
	}
)

var (
	config Config
	once   sync.Once
)

// Get returns config.
func Get() *Config {
	once.Do(func() {
		err := cleanenv.ReadEnv(&config)
		if err != nil {
			log.Fatal("failed to read env", err)
		}
	})

	return &config
}
