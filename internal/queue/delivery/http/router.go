package http

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoutes(r fiber.Router) {
	r.Get("", h.FindQueue())
	r.Post("/enqueue", h.Enqueue())
	r.Post("/dequeue", h.Dequeue())
}
