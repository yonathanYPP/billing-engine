package repositories

import (
	"billing-engine/app/models"

	"gorm.io/gorm"
)

// PaymentRepository interface
type PaymentRepository interface {
	CreatePayment(LoanID uint, totalAmount float64, InstallmentNumber int64) (*models.Payment, error)
	GetPaymentsByLoanID(LoanID uint) ([]models.Payment, int64, error)
}

// PaymentRepositoryImpl struct
type PaymentRepositoryImpl struct {
	db *gorm.DB
}

// NewPaymentRepository membuat instance repository
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}

// CreatePayment menambahkan payment baru ke database
func (u *PaymentRepositoryImpl) CreatePayment(LoanID uint, totalAmount float64, InstallmentNumber int64) (*models.Payment, error) {
	Payment := &models.Payment{
		LoanID:            LoanID,
		TotalAmount:       totalAmount,
		InstallmentNumber: InstallmentNumber,
	}
	err := u.db.Create(Payment).Error
	return Payment, err
}

func (u *PaymentRepositoryImpl) GetPaymentsByLoanID(LoanID uint) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64

	// Query: Ambil data dengan order DESC dan hitung total
	result := u.db.Where("loan_id = ?", LoanID).
		Order("created_at DESC").
		Find(&payments).
		Count(&count)

	if result.Error != nil {
		return payments, 0, result.Error
	}

	return payments, count, nil
}
