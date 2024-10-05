package repository

import (
	"gin-project/config"
	"gin-project/models"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the plain password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares the provided password with the hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// FindUserByUsername finds a user by their username (from MySQL)
func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetAllUsers fetches all users from the database.
func GetAllUsers(users *[]models.User) error {
	if result := config.DB.Find(users); result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateUser membuat pengguna baru dengan saldo awal
func CreateUser(username, password string) (models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: username,
		Password: hashedPassword,
		Balance:  1000.0, // Saldo awal
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
