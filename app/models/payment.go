package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	LoanID            uint           `gorm:"index" json:"loan_id"`
	TotalAmount       float64        `json:"total_amount"`
	InstallmentNumber int64          `json:"installment_number"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
