package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	SenderID   uint    `json:"sender_id"`
	ReceiverID uint    `json:"receiver_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"` // e.g., "completed", "pending", "failed"
}
