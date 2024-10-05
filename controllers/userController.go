package controllers

import (
	"gin-project/models"
	"gin-project/repository"
	"gin-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new user using repository and hash the password
	user, err := repository.CreateUser(input.Username, input.Password)
	if err != nil {
		utils.RespondError(c, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "User created successfully", user)
}

func GetUsers(c *gin.Context) {
	var users []models.User // Kita gunakan struct User dari repository
	err := repository.GetAllUsers(&users)

	if err != nil {
		utils.RespondError(c, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Users fetched successfully", users)
}

func GetUserHistory(c *gin.Context) {
	// Ambil ID pengguna yang saat ini login dari context
	user, _ := c.Get("user")

	userID := user.(models.User).ID

	// Ambil log untuk pengguna
	logs, err := repository.GetUserLogs(userID)
	if err != nil {
		utils.RespondError(c, "Failed to fetch logs", http.StatusInternalServerError)
		return
	}

	utils.RespondSuccess(c, "Logs fetched successfully", logs)
}
