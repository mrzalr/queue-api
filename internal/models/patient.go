package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name      string    `json:"name"`
	DoB       time.Time `json:"date_of_birth"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Patient) BeforeCreate(db *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
