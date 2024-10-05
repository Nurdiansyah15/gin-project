package models

import (
	"time"

	"gorm.io/gorm"
)

type UserHistory struct {
	gorm.Model
	UserID    uint      `json:"user_id"`   // ID pengguna yang melakukan aktivitas
	Action    string    `json:"action"`    // Jenis aktivitas (misalnya login, logout, transfer)
	Timestamp time.Time `json:"timestamp"` // Waktu aktivitas
}
