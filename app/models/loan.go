package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	TotalAmount       float64        `json:"total_amount"`
	InstallmentAmount float64        `json:"installment_amount"`
	InstallmentType   string         `json:"installment_type"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
