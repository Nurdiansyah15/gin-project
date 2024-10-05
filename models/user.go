package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"unique"`
	Password string  `json:"-"`       // Hashed password
	Balance  float64 `json:"balance"` // Saldo pengguna
}
