package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"goraft/internal/adapters/config"
)

func main() {
	// โหลด config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// สตาร์ท Echo
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, world!")
	})

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting Echo server at %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatalf("Echo error: %v", err)
	}
}
