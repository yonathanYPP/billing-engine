package services

import (
	"billing-engine/app/usecases"
	"billing-engine/app/viewmodels"
	"errors"
	"math"
	"time"

	"gorm.io/gorm"
)

// LoanService interface
type LoanService interface {
	GetLoanByID(loanID uint) (*viewmodels.LoanViewmodel, error)
	CreateWeeklyLoan(totalAmount float64, weekNumber int) (*viewmodels.LoanViewmodel, error)
	GetPendingWeeklyPayment(loanID uint) (result *viewmodels.PendingPaymentViewmodel, err error)
	MakePayment(loanID uint, totalPaidAmount float64) error
}

// LoanServiceImpl struct
type LoanServiceImpl struct {
	loanUsecase    usecases.LoanUsecase
	paymentService PaymentService
}

// NewLoanService membuat instance service
func NewLoanService(db *gorm.DB) LoanService {
	loanUsecase := usecases.NewLoanUsecase(db)
	paymentService := NewPaymentService(db)
	return &LoanServiceImpl{
		loanUsecase:    loanUsecase,
		paymentService: paymentService,
	}
}

// GetLoanByID memanggil usecase
func (s *LoanServiceImpl) GetLoanByID(loanID uint) (*viewmodels.LoanViewmodel, error) {
	loan, err := s.loanUsecase.GetLoanByID(loanID)
	if err != nil {
		return nil, err
	}

	res := &viewmodels.LoanViewmodel{
		ID:                loan.ID,
		TotalAmount:       loan.TotalAmount,
		InstallmentAmount: loan.InstallmentAmount,
		InstallmentType:   loan.InstallmentType,
		CreatedAt:         loan.CreatedAt,
		UpdatedAt:         loan.UpdatedAt,
		DeletedAt:         loan.DeletedAt.Time,
	}
	return res, err
}

// CreateLoan service untuk create loan
func (s *LoanServiceImpl) CreateLoan(totalAmount, InstallmentAmount float64, InstallmentType string) (*viewmodels.LoanViewmodel, error) {
	if totalAmount <= 0 || InstallmentAmount <= 0 {
		return nil, errors.New("invalid loan details")
	}

	loan, err := s.loanUsecase.CreateLoan(totalAmount, InstallmentAmount, InstallmentType)
	if err != nil {
		return nil, err
	}

	res := &viewmodels.LoanViewmodel{
		ID:                loan.ID,
		TotalAmount:       loan.TotalAmount,
		InstallmentAmount: loan.InstallmentAmount,
		InstallmentType:   loan.InstallmentType,
		CreatedAt:         loan.CreatedAt,
		UpdatedAt:         loan.UpdatedAt,
		DeletedAt:         loan.DeletedAt.Time,
	}

	return res, err
}

// CreateWeeklyLoan menyimpan data loan type weekly
func (s *LoanServiceImpl) CreateWeeklyLoan(totalAmount float64, weekNumber int) (*viewmodels.LoanViewmodel, error) {
	if totalAmount <= 0 || weekNumber <= 0 {
		return nil, errors.New("invalid loan details")
	}

	outstandingAmount := totalAmount + (totalAmount * 0.10)
	InstallmentAmount := outstandingAmount / float64(weekNumber)

	loan, err := s.CreateLoan(outstandingAmount, InstallmentAmount, "weekly")
	return loan, err
}

// MakePayment memproses pembayaran
func (s *LoanServiceImpl) MakePayment(loanID uint, totalPaidAmount float64) error {
	var deliquentWeek int64
	loan, err := s.loanUsecase.GetLoanByID(loanID)
	if err != nil {
		return err
	}

	_, count, err := s.paymentService.GetPaymentsByLoanID(loanID)
	if err != nil {
		return err
	}

	// Hitung minggu yang sedang berjalan
	now := time.Now()
	weeksPassed := int64(now.Sub(loan.CreatedAt).Hours() / 24 / 7)
	if weeksPassed >= 50 {
		weeksPassed = 50
		deliquentWeek = weeksPassed - count
	} else {
		weeksPassed += 1                          // dihitung termasuk pembayaran minggu ini yang belum jatuh tempo
		deliquentWeek = (weeksPassed - count) - 1 // dihitung tidak termasuk pembayaran minggu ini yang belum jatuh tempo, hanya minggu yang sudah lewat jatuh tempo
	}

	pendingPaymentWeek := weeksPassed - count

	totalBill := loan.InstallmentAmount * float64(pendingPaymentWeek)
	totalPendingBill := loan.InstallmentAmount * float64(deliquentWeek)

	if totalPaidAmount != totalPendingBill && totalPaidAmount != totalBill { // jika hanya bayar yang lewat jatuh tempo
		return errors.New("Payment amount is insufficient.")
	}

	paymentWeek := deliquentWeek
	if totalPaidAmount == totalBill { // jika juga bayar yang jatuh tempo di minggu ini
		paymentWeek = pendingPaymentWeek
	}

	if paymentWeek == 0 {
		return errors.New("Payment amount is insufficient.")
	}

	_, err = s.paymentService.PayPendingPayment(loanID, loan.InstallmentAmount, count, paymentWeek)

	return nil
}

// IsDelinquent menghitung saldo yang menunggak dan jadwal pembayaran
func (s *LoanServiceImpl) GetPendingWeeklyPayment(loanID uint) (result *viewmodels.PendingPaymentViewmodel, err error) {
	var deliquentWeek int64
	loan, err := s.loanUsecase.GetLoanByID(loanID)
	if err != nil {
		return result, err
	}

	_, count, err := s.paymentService.GetPaymentsByLoanID(loanID)
	if err != nil {
		return result, err
	}

	// Hitung minggu berjalan
	now := time.Now()
	weeksPassed := int64(now.Sub(loan.CreatedAt).Hours() / 24 / 7)
	if weeksPassed >= 50 {
		weeksPassed = 50
		deliquentWeek = int64(math.Max(0, float64(weeksPassed-count)))
	} else {
		weeksPassed += 1                                                   // dihitung termasuk pembayaran minggu ini yang belum jatuh tempo
		deliquentWeek = int64(math.Max(0, float64((weeksPassed-count)-1))) // dihitung tidak termasuk pembayaran minggu ini yang belum jatuh tempo, hanya minggu yang sudah lewat jatuh tempo
	}

	pendingPaymentWeek := weeksPassed - count
	IsDelinquent := deliquentWeek > 1
	paidInstallmentTotal := loan.InstallmentAmount * float64(count)

	var pendingPayments []viewmodels.PaymentViewmodel

	for i := pendingPaymentWeek; i > int64(0); i-- {
		InstallmentNumber := count + i
		// Hitung due date berdasarkan minggu terakhir pinjaman
		dueDate := loan.CreatedAt.Add(time.Duration(InstallmentNumber) * 7 * 24 * time.Hour)
		pendingPayments = append(pendingPayments, viewmodels.PaymentViewmodel{
			TotalAmount:       loan.InstallmentAmount,
			InstallmentNumber: InstallmentNumber,
			DueDate:           dueDate,
		})
	}

	result = &viewmodels.PendingPaymentViewmodel{
		IsDelinquent:           IsDelinquent,
		OutStandingTotal:       loan.TotalAmount - paidInstallmentTotal,
		PaidInstallmentTotal:   paidInstallmentTotal,
		PendingPaymentTotal:    loan.InstallmentAmount * float64(pendingPaymentWeek),
		DelinquentPaymentTotal: loan.InstallmentAmount * float64(deliquentWeek),
		PendingPayment:         pendingPayments,
	}

	return result, nil
}
