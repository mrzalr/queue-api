package server

import (
	"github.com/gofiber/fiber/v2"
	patientHttp "github.com/mrzalr/queue-api/internal/patient/delivery/http"
	patientRepo "github.com/mrzalr/queue-api/internal/patient/repository/mysql"
	patientUcase "github.com/mrzalr/queue-api/internal/patient/usecase"
	queueHttp "github.com/mrzalr/queue-api/internal/queue/delivery/http"
	queueRepo "github.com/mrzalr/queue-api/internal/queue/repository/mysql"
	queueUcase "github.com/mrzalr/queue-api/internal/queue/usecase"
)

func (s *server) MapHandlers(app *fiber.App) {
	patientRepository := patientRepo.New(s.DB)
	queueRepository := queueRepo.New(s.DB)

	patientUsecase := patientUcase.New(patientRepository)
	queueUsecase := queueUcase.New(queueRepository)

	patientHandler := patientHttp.New(patientUsecase)
	queueHandler := queueHttp.New(queueUsecase)

	v1 := app.Group("v1")

	patient := v1.Group("/patient")
	patientHandler.MapRoutes(patient)

	queue := v1.Group("/queue")
	queueHandler.MapRoutes(queue)
}
