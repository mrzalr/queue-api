package server

import (
	"github.com/gofiber/fiber/v2"
	patientHttp "github.com/mrzalr/queue-api/internal/patient/delivery/http"
	patientRepo "github.com/mrzalr/queue-api/internal/patient/repository/mysql"
	patientUcase "github.com/mrzalr/queue-api/internal/patient/usecase"
)

func (s *server) MapHandlers(app *fiber.App) {
	patientRepository := patientRepo.New(s.DB)

	patientUsecase := patientUcase.New(patientRepository)

	patientHandler := patientHttp.New(patientUsecase)

	v1 := app.Group("v1")

	patient := v1.Group("/patient")
	patientHandler.MapRoutes(patient)
}
