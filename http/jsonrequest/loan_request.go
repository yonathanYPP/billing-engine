package jsonrequest

// LoanRequest struct
type LoanRequest struct {
	TotalAmount float64 `json:"total_amount"`
	WeekNumber  int     `json:"week_number"`
}

// InstallmentPaymentRequest struct
type InstallmentPaymentRequest struct {
	LoanID          uint64  `json:"loan_id"`
	TotalPaidAmount float64 `json:"total_paid_amount"`
}
