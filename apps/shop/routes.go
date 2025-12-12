package shop

import (
	"github.com/gin-gonic/gin"

	"github.com/farhapartex/dealer-marketplace-be/apps/shop/services"
	"github.com/farhapartex/dealer-marketplace-be/apps/shop/views"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/middleware"
)

func RegisterRoutes(router *gin.RouterGroup) {
	shopService := services.NewShopService()
	shopView := views.NewShopView(shopService)

	categoryService := services.NewCategoryService()
	categoryView := views.NewCategoryView(categoryService)

	shopGroup := router.Group("/shops")
	shopGroup.Use(middleware.AuthMiddleware())
	{
		shopGroup.POST("", shopView.CreateShopHandler)
	}

	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("", categoryView.GetCategoriesHandler)
	}
}
