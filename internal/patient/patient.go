package patient

import (
	"github.com/google/uuid"
	"github.com/mrzalr/queue-api/internal/models"
)

type Repository interface {
	Find() ([]models.Patient, error)
	FindByID(patientID uuid.UUID) (models.Patient, error)
	Create(patient models.Patient) (models.Patient, error)
	Update(patient models.Patient) (models.Patient, error)
	Delete(patientID uuid.UUID) error
}

type Usecase interface {
	Find() ([]models.Patient, error)
	FindByID(patientID uuid.UUID) (models.Patient, error)
	Create(patient models.Patient) (models.Patient, error)
	Update(patient models.Patient) (models.Patient, error)
	Delete(patientID uuid.UUID) error
}
