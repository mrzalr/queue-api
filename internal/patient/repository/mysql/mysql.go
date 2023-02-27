package mysql

import (
	"github.com/google/uuid"
	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/patient"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) patient.Repository {
	return &repository{db}
}

func (r *repository) Find() ([]models.Patient, error) {
	patients := []models.Patient{}
	tx := r.db.Find(&patients)
	return patients, tx.Error
}

func (r *repository) FindByID(patientID uuid.UUID) (models.Patient, error) {
	patient := models.Patient{}
	tx := r.db.Where("id = ?", patientID).First(&patient)
	return patient, tx.Error
}

func (r *repository) Create(patient models.Patient) (models.Patient, error) {
	tx := r.db.Create(&patient)
	return patient, tx.Error
}

func (r *repository) Update(patient models.Patient) (models.Patient, error) {
	tx := r.db.Where("id = ?", patient.ID).Updates(&patient)
	return patient, tx.Error
}

func (r *repository) Delete(patientID uuid.UUID) error {
	tx := r.db.Where("id = ?", patientID).Delete(&models.Patient{})
	return tx.Error
}
