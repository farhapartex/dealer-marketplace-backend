package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/farhapartex/dealer-marketplace-be/apps/shop/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/shop/services"
)

type ShopView struct {
	shopService *services.ShopService
}

func NewShopView(shopService *services.ShopService) *ShopView {
	return &ShopView{
		shopService: shopService,
	}
}

// CreateShopHandler handles the creation of a new shop
func (v *ShopView) CreateShopHandler(c *gin.Context) {
	// Extract user_id from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User ID not found in context",
		})
		return
	}

	// Type assertion to uuid.UUID
	uid, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Invalid user ID format",
		})
		return
	}

	// Bind request payload
	var req dtos.CreateShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Call service to create shop
	response, err := v.shopService.CreateShop(uid, &req)
	if err != nil {
		if err.Error() == "user already has a shop" {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": response,
	})
}
