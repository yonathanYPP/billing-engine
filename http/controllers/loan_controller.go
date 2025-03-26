package controllers

import (
	services "billing-engine/app/service"
	"billing-engine/http/jsonrequest"
	"billing-engine/http/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoanController struct
type LoanController struct {
	loanService services.LoanService
}

// NewLoanController membuat instance controller
func NewLoanController(db *gorm.DB) *LoanController {
	loanService := services.NewLoanService(db)

	return &LoanController{
		loanService: loanService,
	}
}

// GetLoanByID handler
func (lc *LoanController) GetLoanByID(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loan, err := lc.loanService.GetLoanByID(uint(loanID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	response := views.NewLoanResponse(loan)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

// GetPendingWeeklyPayment handler
func (lc *LoanController) GetPendingWeeklyPayment(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	res, err := lc.loanService.GetPendingWeeklyPayment(uint(loanID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	response := views.NewPendingPaymentResponse(res)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

// CreateLoan membuat pinjaman baru mingguan
func (lc *LoanController) CreateWeeklyLoan(c *gin.Context) {
	var loanReq jsonrequest.LoanRequest
	if err := c.ShouldBindJSON(&loanReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := lc.loanService.CreateWeeklyLoan(loanReq.TotalAmount, loanReq.WeekNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := views.NewLoanResponse(loan)
	c.JSON(http.StatusCreated, gin.H{"message": "Loan created successfully", "data": response})
}

// MakePayment membuat pembayaran baru
func (lc *LoanController) MakePayment(c *gin.Context) {
	var installmentPayment jsonrequest.InstallmentPaymentRequest
	if err := c.ShouldBindJSON(&installmentPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := lc.loanService.MakePayment(uint(installmentPayment.LoanID), installmentPayment.TotalPaidAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully"})
}
