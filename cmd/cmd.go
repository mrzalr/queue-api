package cmd

import (
	"log"

	"github.com/mrzalr/queue-api/config"
	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/server"
	"github.com/mrzalr/queue-api/pkg/db/mysql"
)

func StartApplication() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error when loading .env | err : %s", err)
	}

	db, err := mysql.New()
	if err != nil {
		log.Fatalf("Error when create new mysql connection | err : %s", err)
	}
	db.AutoMigrate(&models.Patient{}, &models.Queue{})

	s := server.New(db)
	if err := s.Run(); err != nil {
		log.Fatalf("Error when run a server | err : %s", err)
	}

}
