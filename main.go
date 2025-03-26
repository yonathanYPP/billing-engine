package main

import (
	"billing-engine/database"
	"billing-engine/http/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.InitDB()

	loanController := controllers.NewLoanController(db)

	r.GET("/loan/:id", loanController.GetLoanByID)
	r.GET("/loan/:id/pending-weekly-payment", loanController.GetPendingWeeklyPayment)
	r.POST("/loan", loanController.CreateWeeklyLoan)
	r.POST("/make-payment", loanController.MakePayment)

	r.Run(":8080")
}
