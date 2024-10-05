package controllers

import (
	"gin-project/models"
	"gin-project/repository"
	"gin-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new user with hashed password
	user, err := repository.CreateUser(input.Username, input.Password)
	if err != nil {
		utils.RespondError(c, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "User registered successfully", user)
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the user by username
	user, err := repository.FindUserByUsername(input.Username)
	if err != nil {
		utils.RespondError(c, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Validate the password
	if !repository.CheckPasswordHash(input.Password, user.Password) {
		utils.RespondError(c, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(input.Username)
	if err != nil {
		utils.RespondError(c, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Log user login
	err = repository.CreateLog(user.ID, "login") // Menambahkan log login
	if err != nil {
		utils.RespondError(c, "Failed to log user activity", http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Login successful", gin.H{"token": token})
}

func Logout(c *gin.Context) {
	// Mendapatkan token dari header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.RespondError(c, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	tokenString := strings.Split(authHeader, "Bearer ")[1]

	// Tambahkan token ke blacklist
	err := repository.AddToBlacklist(tokenString)
	if err != nil {
		utils.RespondError(c, "Failed to blacklist token", http.StatusInternalServerError)
		return
	}

	user, _ := c.Get("user")
	userID := user.(models.User).ID

	// Log user logout
	errLog := repository.CreateLog(userID, "logout") // Menambahkan log logout
	if errLog != nil {
		utils.RespondError(c, "Failed to log user activity", http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Logout successful", nil)
}
