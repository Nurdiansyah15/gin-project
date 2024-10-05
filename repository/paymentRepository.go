package repository

import (
	"errors"
	"gin-project/config"
	"gin-project/models"
)

// CreatePayment membuat pembayaran baru dalam database
func CreatePayment(senderID uint, receiverID uint, amount float64) (models.Payment, error) {
	// Cek saldo pengirim
	var sender models.User
	if err := config.DB.First(&sender, senderID).Error; err != nil {
		return models.Payment{}, err
	}

	if sender.Balance < amount {
		return models.Payment{}, errors.New("insufficient balance")
	}

	// Update saldo pengirim
	sender.Balance -= amount
	if err := config.DB.Save(&sender).Error; err != nil {
		return models.Payment{}, err
	}

	// Update saldo penerima
	var receiver models.User
	if err := config.DB.First(&receiver, receiverID).Error; err != nil {
		return models.Payment{}, err
	}

	receiver.Balance += amount
	if err := config.DB.Save(&receiver).Error; err != nil {
		return models.Payment{}, err
	}

	// Buat record pembayaran
	payment := models.Payment{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     amount,
		Status:     "completed", // Pembayaran berhasil
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

// GetPaymentsByUser mengembalikan semua pembayaran yang terkait dengan user
func GetPaymentsByUser(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := config.DB.Where("sender_id = ? OR receiver_id = ?", userID, userID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
