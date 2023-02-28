package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/patient"
)

type handler struct {
	usecase patient.Usecase
}

func New(usecase patient.Usecase) *handler {
	return &handler{usecase}
}

func (h *handler) Find() fiber.Handler {
	return func(c *fiber.Ctx) error {
		patients, err := h.usecase.Find()
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusOK).
			JSON(models.ResponseOK(patients))
	}
}

func (h *handler) FindByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		patientID, err := uuid.Parse(c.Params("id"))
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
		}

		patient, err := h.usecase.FindByID(patientID)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusOK).
			JSON(models.ResponseOK(patient))
	}
}

func (h *handler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		patient := models.Patient{}
		err := c.BodyParser(&patient)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
		}

		patient, err = h.usecase.Create(patient)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusCreated).
			JSON(models.ResponseCreated(patient))
	}
}

func (h *handler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		patientID, err := uuid.Parse(c.Params("id"))
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
		}

		patientUpdate := models.Patient{ID: patientID}
		err = c.BodyParser(&patientUpdate)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
		}

		patient, err := h.usecase.Update(patientUpdate)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusOK).
			JSON(models.ResponseOK(patient))
	}
}

func (h *handler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		patientID, err := uuid.Parse(c.Params("id"))
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
		}

		err = h.usecase.Delete(patientID)
		if err != nil {
			errs := []string{err.Error()}
			return c.Status(http.StatusBadGateway).
				JSON(models.ResponseBadGateway(errs))
		}

		return c.Status(http.StatusOK).
			JSON(models.ResponseOK(nil))
	}
}
