package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type server struct {
	DB  *gorm.DB
	App *fiber.App
}

func New(db *gorm.DB) *server {
	return &server{
		DB:  db,
		App: fiber.New(),
	}
}

func (s *server) Run() error {
	s.MapHandlers(s.App)

	port := os.Getenv("APP_PORT")
	log.Printf("Server is running on port %s", port)
	return s.App.Listen(fmt.Sprintf(":%s", port))
}
