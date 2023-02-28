package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/queue-api/internal/queue"
)

type handler struct {
	usecase queue.Usecase
}

func New(usecase queue.Usecase) *handler {
	return &handler{usecase}
}

func (h *handler) Find() fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}
