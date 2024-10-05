package repository

import (
	"gin-project/config"
	"gin-project/models"

	"gorm.io/gorm"
)

// AddToBlacklist akan menambahkan token ke daftar hitam
func AddToBlacklist(token string) error {
	blacklistedToken := models.BlacklistedToken{
		Token: token,
	}
	if err := config.DB.Create(&blacklistedToken).Error; err != nil {
		return err
	}
	return nil
}

// IsTokenBlacklisted akan mengecek apakah token sudah di-blacklist
func IsTokenBlacklisted(token string) (bool, error) {
	var blacklistedToken models.BlacklistedToken
	err := config.DB.Where("token = ?", token).First(&blacklistedToken).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // Token tidak di-blacklist
		}
		return false, err
	}
	return true, nil // Token ditemukan, berarti sudah di-blacklist
}
