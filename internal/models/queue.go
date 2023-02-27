package models

import "time"

type Queue struct {
	Number    int       `json:"queue_number"`
	IsDone    bool      `json:"is_done"`
	PatientID int       `json:"patient_id"`
	Patient   Patient   `json:"patient"`
	CreatedAt time.Time `json:"created_at"`
}
