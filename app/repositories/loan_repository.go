package repositories

import (
	"billing-engine/app/models"

	"gorm.io/gorm"
)

// LoanRepository interface
type LoanRepository interface {
	GetLoanByID(loanID uint) (*models.Loan, error)
	CreateLoan(totalAmount, InstallmentAmount float64, InstallmentType string) (*models.Loan, error)
	UpdateLoan(loanID uint, totalAmount, InstallmentAmount float64, InstallmentType string) error
}

// LoanRepositoryImpl struct
type LoanRepositoryImpl struct {
	db *gorm.DB
}

// NewLoanRepository membuat instance repository
func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &LoanRepositoryImpl{db: db}
}

// GetLoanByID mengambil data berdasarkan ID
func (r *LoanRepositoryImpl) GetLoanByID(loanID uint) (*models.Loan, error) {
	var loan models.Loan
	result := r.db.First(&loan, loanID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &loan, nil
}

// CreateLoan menambahkan pinjaman baru ke database
func (u *LoanRepositoryImpl) CreateLoan(totalAmount, InstallmentAmount float64, InstallmentType string) (*models.Loan, error) {
	loan := &models.Loan{
		TotalAmount:       totalAmount,
		InstallmentAmount: InstallmentAmount,
		InstallmentType:   InstallmentType,
	}
	err := u.db.Create(loan).Error
	return loan, err
}

// UpdateLoan memperbarui data loan
func (u *LoanRepositoryImpl) UpdateLoan(loanID uint, totalAmount, InstallmentAmount float64, InstallmentType string) error {
	loan := &models.Loan{
		ID:                loanID,
		TotalAmount:       totalAmount,
		InstallmentAmount: InstallmentAmount,
		InstallmentType:   InstallmentType,
	}
	return u.db.Save(loan).Error
}
