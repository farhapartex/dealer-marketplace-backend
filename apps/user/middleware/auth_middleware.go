package middleware

import (
	"net/http"
	"strings"

	"github.com/farhapartex/dealer-marketplace-be/apps/user/models"
	"github.com/farhapartex/dealer-marketplace-be/pkg/database"
	"github.com/farhapartex/dealer-marketplace-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid authorization header format. Expected 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]
		userID, err := utils.VerifyJWT(token)
		if err != nil || userID == uuid.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		var user models.User
		if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "User not found",
			})
			c.Abort()
			return
		}

		if !user.IsActive {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "User account is not active",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("user", user)
		c.Next()
	}
}
