package usecases

import (
	"billing-engine/app/models"
	"billing-engine/app/repositories"

	"gorm.io/gorm"
)

// PaymentUsecase interface
type PaymentUsecase interface {
	CreatePayment(LoanID uint, totalAmount float64, InstallmentNumber int64) (*models.Payment, error)
	GetPaymentsByLoanID(LoanID uint) ([]models.Payment, int64, error)
}

// PaymentUsecaseImpl struct
type PaymentUsecaseImpl struct {
	PaymentRepo repositories.PaymentRepository
}

// NewPaymentUsecase membuat instance usecase
func NewPaymentUsecase(db *gorm.DB) PaymentUsecase {
	PaymentRepo := repositories.NewPaymentRepository(db)
	return &PaymentUsecaseImpl{PaymentRepo: PaymentRepo}
}

// CreatePayment menambahkan payment baru ke database
func (u *PaymentUsecaseImpl) CreatePayment(LoanID uint, totalAmount float64, InstallmentNumber int64) (*models.Payment, error) {
	Payment, err := u.PaymentRepo.CreatePayment(LoanID, totalAmount, InstallmentNumber)
	return Payment, err
}

func (u *PaymentUsecaseImpl) GetPaymentsByLoanID(LoanID uint) ([]models.Payment, int64, error) {
	payments, count, err := u.PaymentRepo.GetPaymentsByLoanID(LoanID)
	return payments, count, err
}
