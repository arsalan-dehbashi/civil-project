package app

import (
	"github.com/arsalan-dehbashi/civil-project.git/internal/config"
	"github.com/arsalan-dehbashi/civil-project.git/internal/db"
	"github.com/arsalan-dehbashi/civil-project.git/pkg/redis"
	"gorm.io/gorm"
)

type App struct {
	Cfg   *config.Config
	DB    *gorm.DB
	Redis *redis.Client
}

func New(cfg *config.Config) (*App, error) {
	gormDB, err := db.Init(cfg)
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Cfg:   cfg,
		DB:    gormDB,
		Redis: redisClient,
	}, nil
}
