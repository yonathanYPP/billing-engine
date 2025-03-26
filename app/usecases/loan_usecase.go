package usecases

import (
	"billing-engine/app/models"
	"billing-engine/app/repositories"

	"gorm.io/gorm"
)

// LoanUsecase interface
type LoanUsecase interface {
	GetLoanByID(loanID uint) (*models.Loan, error)
	CreateLoan(totalAmount, InstallmentAmount float64, InstallmentType string) (*models.Loan, error)
	UpdateLoan(loanID uint, totalAmount, InstallmentAmount float64, InstallmentType string) error
}

// LoanUsecaseImpl struct
type LoanUsecaseImpl struct {
	loanRepo repositories.LoanRepository
}

// NewLoanUsecase membuat instance usecase
func NewLoanUsecase(db *gorm.DB) LoanUsecase {
	loanRepo := repositories.NewLoanRepository(db)
	return &LoanUsecaseImpl{loanRepo: loanRepo}
}

// GetLoanByID memanggil repository
func (u *LoanUsecaseImpl) GetLoanByID(loanID uint) (*models.Loan, error) {
	return u.loanRepo.GetLoanByID(loanID)
}

// CreateLoan menambahkan pinjaman baru ke database
func (u *LoanUsecaseImpl) CreateLoan(totalAmount, InstallmentAmount float64, InstallmentType string) (*models.Loan, error) {
	loan, err := u.loanRepo.CreateLoan(totalAmount, InstallmentAmount, InstallmentType)
	return loan, err
}

// UpdateLoan update pinjaman baru ke database
func (u *LoanUsecaseImpl) UpdateLoan(loanID uint, totalAmount, InstallmentAmount float64, InstallmentType string) error {
	err := u.loanRepo.UpdateLoan(loanID, totalAmount, InstallmentAmount, InstallmentType)
	return err
}
