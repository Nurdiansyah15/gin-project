package controllers

import (
	"gin-project/models"
	"gin-project/repository"
	"gin-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentInput struct {
	ReceiverID uint    `json:"receiver_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

func CreatePayment(c *gin.Context) {
	var input PaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	// Ambil ID pengguna yang saat ini login dari context
	user, _ := c.Get("user")

	senderID := user.(models.User).ID

	// Buat pembayaran
	payment, err := repository.CreatePayment(senderID, input.ReceiverID, input.Amount)
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log user create payment
	err = repository.CreateLog(senderID, "create payment")
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// log user receive payment
	err = repository.CreateLog(input.ReceiverID, "receive payment")
	if err != nil {
		utils.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Payment created successfully", payment)
}

func GetUserPayments(c *gin.Context) {
	// Ambil ID pengguna yang saat ini login dari context
	user, _ := c.Get("user")

	userID := user.(models.User).ID

	payments, err := repository.GetPaymentsByUser(userID)
	if err != nil {
		utils.RespondError(c, "Failed to fetch payments", http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Payments fetched successfully", payments)
}
