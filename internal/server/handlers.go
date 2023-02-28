package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/queue-api/internal/patient/delivery/http"
	"github.com/mrzalr/queue-api/internal/patient/repository/mysql"
	"github.com/mrzalr/queue-api/internal/patient/usecase"
)

func (s *server) MapHandlers(app *fiber.App) {
	patientRepository := mysql.New(s.DB)

	patientUsecase := usecase.New(patientRepository)

	patientHandler := http.New(patientUsecase)

	v1 := app.Group("v1")

	patient := v1.Group("/patient")
	patientHandler.MapRoutes(patient)
}
