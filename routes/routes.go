package routes

import (
	"gin-project/controllers"
	"gin-project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public route
	r.POST("/login", controllers.Login)

	r.POST("/register", controllers.Register)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/users", controllers.GetUsers)
	auth.POST("/users", controllers.CreateUser)
	auth.GET("/history", controllers.GetUserHistory)

	auth.POST("/logout", controllers.Logout)

	// Payment routes
	auth.POST("/payments", controllers.CreatePayment)  // Endpoint untuk membuat pembayaran
	auth.GET("/payments", controllers.GetUserPayments) // Endpoint untuk mengambil riwayat pembayaran
}
