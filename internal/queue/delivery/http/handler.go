package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/queue"
)

type handler struct {
	usecase queue.Usecase
}

func New(usecase queue.Usecase) *handler {
	return &handler{usecase}
}

func (h *handler) FindQueue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		queue, err := h.usecase.FindQueue()
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusOK).
			JSON(models.ResponseOK(queue))
	}
}

func (h *handler) Dequeue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		queue, err := h.usecase.Dequeue()
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusOK).
			JSON(models.ResponseOK(queue))
	}
}

func (h *handler) Enqueue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		newQueue := models.Queue{}
		err := c.BodyParser(&newQueue)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
		}

		queue, err := h.usecase.Enqueue(newQueue)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusOK).
			JSON(models.ResponseOK(queue))
	}
}
