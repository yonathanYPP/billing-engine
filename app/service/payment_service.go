package services

import (
	"billing-engine/app/models"
	"billing-engine/app/usecases"
	"sync"

	"gorm.io/gorm"
)

// PaymentService interface
type PaymentService interface {
	PayPendingPayment(LoanID uint, totalAmount float64, startInstallment, pendingInstallment int64) ([]*models.Payment, error)
	GetPaymentsByLoanID(LoanID uint) ([]models.Payment, int64, error)
}

// PaymentServiceImpl struct
type PaymentServiceImpl struct {
	PaymentUsecase usecases.PaymentUsecase
}

// NewPaymentService membuat instance service
func NewPaymentService(db *gorm.DB) PaymentService {
	PaymentUsecase := usecases.NewPaymentUsecase(db)
	return &PaymentServiceImpl{PaymentUsecase: PaymentUsecase}
}

// CreatePayment menangani validasi sebelum menyimpan Payment
func (s *PaymentServiceImpl) PayPendingPayment(LoanID uint, totalAmount float64, startInstallment, pendingInstallment int64) ([]*models.Payment, error) {
	var wg sync.WaitGroup
	ch := make(chan int64, pendingInstallment) // Channel untuk mengontrol urutan
	errCh := make(chan error, pendingInstallment)
	// Slice untuk menyimpan hasil insert
	var paymentsInserted []*models.Payment

	for i := int64(1); i <= pendingInstallment; i++ {
		wg.Add(1)
		go func(Installment int64) {
			defer wg.Done()

			// Masukkan nomor Installment ke channel (untuk menjaga urutan)
			ch <- Installment

			payment, err := s.PaymentUsecase.CreatePayment(LoanID, totalAmount, Installment)
			if err != nil {
				errCh <- err
				return
			}

			paymentsInserted = append(paymentsInserted, payment)
		}(startInstallment + i)
	}

	wg.Wait()
	close(ch)
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return paymentsInserted, nil
}

// CreatePayment menambahkan pinjaman baru ke database
func (s *PaymentServiceImpl) GetPaymentsByLoanID(LoanID uint) ([]models.Payment, int64, error) {
	payments, count, err := s.PaymentUsecase.GetPaymentsByLoanID(LoanID)
	return payments, count, err
}
