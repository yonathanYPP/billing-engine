package views

import (
	"billing-engine/app/viewmodels"
)

// LoanResponse struct
type LoanResponse struct {
	ID                uint    `gorm:"primaryKey" json:"id"`
	TotalAmount       float64 `json:"total_amount"`
	InstallmentAmount float64 `json:"installment_amount"`
	InstallmentType   string  `json:"installment_type"`
}

// PendingPaymentResponse struct
type PendingPaymentResponse struct {
	IsDelinquent           bool    `json:"is_delinquent"`
	OutStandingTotal       float64 `json:"outstanding_total"`
	PaidInstallmentTotal   float64 `json:"paid_installment_total"`
	PendingPaymentTotal    float64 `json:"pending_payment_total"`
	DelinquentPaymentTotal float64 `json:"delinquent_payment_total"`
	PendingPayment         []viewmodels.PaymentViewmodel
}

// NewLoanResponse konversi model ke response API
func NewLoanResponse(loan *viewmodels.LoanViewmodel) LoanResponse {
	return LoanResponse{
		ID:                loan.ID,
		TotalAmount:       loan.TotalAmount,
		InstallmentAmount: loan.InstallmentAmount,
		InstallmentType:   loan.InstallmentType,
	}
}

// NewPendingPaymentResponse konversi model ke response API
func NewPendingPaymentResponse(pendingPayment *viewmodels.PendingPaymentViewmodel) PendingPaymentResponse {
	return PendingPaymentResponse{
		IsDelinquent:           pendingPayment.IsDelinquent,
		OutStandingTotal:       pendingPayment.OutStandingTotal,
		PaidInstallmentTotal:   pendingPayment.PaidInstallmentTotal,
		PendingPaymentTotal:    pendingPayment.PendingPaymentTotal,
		DelinquentPaymentTotal: pendingPayment.DelinquentPaymentTotal,
		PendingPayment:         pendingPayment.PendingPayment,
	}
}
