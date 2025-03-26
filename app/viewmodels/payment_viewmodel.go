package viewmodels

import "time"

// PaymentViewmodel struct
type PaymentViewmodel struct {
	TotalAmount       float64   `json:"total_amount"`
	InstallmentNumber int64     `json:"installment_number"`
	DueDate           time.Time `json:"due_date"`
}

// PendingPaymentViewmodel struct
type PendingPaymentViewmodel struct {
	IsDelinquent           bool    `json:"is_delinquent"`
	OutStandingTotal       float64 `json:"outstanding_total"`
	PaidInstallmentTotal   float64 `json:"paid_installment_total"`
	PendingPaymentTotal    float64 `json:"pending_payment_total"`
	DelinquentPaymentTotal float64 `json:"delinquent_payment_total"`
	PendingPayment         []PaymentViewmodel
}
