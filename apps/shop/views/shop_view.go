package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/dealer-marketplace-be/apps/shop/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/shop/services"
	"github.com/farhapartex/dealer-marketplace-be/utils"
)

type ShopView struct {
	shopService *services.ShopService
}

func NewShopView(shopService *services.ShopService) *ShopView {
	return &ShopView{
		shopService: shopService,
	}
}

func (v *ShopView) CreateShopHandler(c *gin.Context) {
	// Extract user_id from context using utility function
	uid, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

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
