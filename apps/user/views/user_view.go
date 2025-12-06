package views

import (
	"github.com/farhapartex/dealer-marketplace-be/apps/user/services"
)

type UserView struct {
	authService *services.AuthService
}

func NewUserView(authService *services.AuthService) *UserView {
	return &UserView{
		authService: authService,
	}
}

// func (v *UserView) CreateUser(c *gin.Context) {
// 	var req dtos.CreateUserRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := v.userService.CreateUser(req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, user)
// }
