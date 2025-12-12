package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/dealer-marketplace-be/apps/shop/services"
)

type CategoryView struct {
	categoryService *services.CategoryService
}

func NewCategoryView(categoryService *services.CategoryService) *CategoryView {
	return &CategoryView{
		categoryService: categoryService,
	}
}

func (v *CategoryView) GetCategoriesHandler(c *gin.Context) {
	categories := v.categoryService.GetAllCategories()

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
