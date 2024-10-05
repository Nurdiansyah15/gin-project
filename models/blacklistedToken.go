package models

import (
	"gorm.io/gorm"
)

// BlacklistedToken struct untuk menyimpan token yang di-blacklist
type BlacklistedToken struct {
	gorm.Model
	Token string `json:"token"`
}
