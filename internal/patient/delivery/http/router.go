package http

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoutes(r fiber.Router) {
	r.Get("", h.Find())
	r.Get("/:id", h.FindByID())
	r.Post("", h.Create())
	r.Put("/:id", h.Update())
	r.Delete("/:id", h.Delete())
}
