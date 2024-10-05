package repository

import (
	"gin-project/config"
	"gin-project/models"
	"time"
)

// CreateLog creates a new log entry in the database
func CreateLog(userID uint, action string) error {
	log := models.UserHistory{
		UserID:    userID,
		Action:    action,
		Timestamp: time.Now(),
	}
	if err := config.DB.Create(&log).Error; err != nil {
		return err
	}
	return nil
}

// GetUserLogs fetches logs for a specific user from the database
func GetUserLogs(userID uint) ([]models.UserHistory, error) {
	var logs []models.UserHistory
	if err := config.DB.Where("user_id = ?", userID).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
