package models

import (
	"time"

	"github.com/google/uuid"
)

type Queue struct {
	Number    int       `json:"queue_number"`
	IsDone    bool      `json:"is_done"`
	PatientID uuid.UUID `json:"patient_id" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	Patient   Patient   `json:"patient"`
}
