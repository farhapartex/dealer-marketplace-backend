package services

import "github.com/farhapartex/dealer-marketplace-be/utils"

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

// GetAllCategories returns all categories with their subcategories
func (s *CategoryService) GetAllCategories() []utils.Category {
	return utils.GetAllCategories()
}

// GetCategoryByName returns a specific category by name
func (s *CategoryService) GetCategoryByName(name string) *utils.Category {
	categories := utils.GetAllCategories()
	for _, category := range categories {
		if category.Name == name {
			return &category
		}
	}
	return nil
}

// GetSubcategories returns all subcategories for a given category
func (s *CategoryService) GetSubcategories(categoryName string) []string {
	return utils.GetSubcategories(categoryName)
}
