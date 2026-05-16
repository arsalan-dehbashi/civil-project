package main

import (
	"log"

	"github.com/arsalan-dehbashi/civil-project.git/internal/app"
	"github.com/arsalan-dehbashi/civil-project.git/internal/config"
	"github.com/arsalan-dehbashi/civil-project.git/internal/http/router"
)

func main() {
	cfg := config.LoadConfig()

	a, err := app.New(cfg)
	if err != nil {
		log.Fatalf("bootstrap failed: %v", err)
	}

	r := router.New(a)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
