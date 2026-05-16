package app

import (
	"github.com/arsalan-dehbashi/civil-project.git/internal/config"
	"github.com/arsalan-dehbashi/civil-project.git/internal/db"
	"gorm.io/gorm"
)

type App struct {
	Cfg *config.Config
	DB  *gorm.DB
}

func New(cfg *config.Config) (*App, error) {
	gormDB, err := db.Init(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Cfg: cfg,
		DB:  gormDB,
	}, nil
}
