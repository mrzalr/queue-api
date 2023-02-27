package usecase

import (
	"github.com/google/uuid"
	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/patient"
)

type usecase struct {
	repository patient.Repository
}

func New(repository patient.Repository) patient.Usecase {
	return &usecase{repository}
}

func (u *usecase) Find() ([]models.Patient, error) {
	return u.repository.Find()
}

func (u *usecase) FindByID(patientID uuid.UUID) (models.Patient, error) {
	return u.repository.FindByID(patientID)
}

func (u *usecase) Create(patient models.Patient) (models.Patient, error) {
	return u.repository.Create(patient)
}

func (u *usecase) Update(patient models.Patient) (models.Patient, error) {
	found, err := u.FindByID(patient.ID)
	if err != nil {
		return models.Patient{}, err
	}

	found.Name = patient.Name
	found.Address = patient.Address
	found.DoB = patient.DoB

	return u.repository.Update(found)
}

func (u *usecase) Delete(patientID uuid.UUID) error {
	_, err := u.FindByID(patientID)
	if err != nil {
		return err
	}

	return u.repository.Delete(patientID)
}
