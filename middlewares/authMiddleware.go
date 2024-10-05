package middlewares

import (
	"gin-project/repository"
	"gin-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondError(c, "Authorization header missing", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Mendapatkan token dari header Authorization
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Cek apakah token sudah di-blacklist
		isBlacklisted, err := repository.IsTokenBlacklisted(tokenString)
		if err != nil {
			utils.RespondError(c, "Failed to validate token", http.StatusInternalServerError)
			c.Abort()
			return
		}

		if isBlacklisted {
			utils.RespondError(c, "Token has been blacklisted", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Validasi token JWT
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			utils.RespondError(c, "Invalid token", http.StatusUnauthorized)
			c.Abort()
			return
		}

		user, err := repository.FindUserByUsername(claims.Username)
		if err != nil {
			utils.RespondError(c, "Failed to fetch user ID", http.StatusInternalServerError)
			c.Abort()
			return
		}

		// Menyimpan username dalam context untuk request berikutnya
		c.Set("user", user)

		c.Next()
	}
}
