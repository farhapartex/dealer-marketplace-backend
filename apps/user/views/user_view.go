package views

import (
	"net/http"

	"github.com/farhapartex/dealer-marketplace-be/apps/user/services"
	"github.com/farhapartex/dealer-marketplace-be/utils"
	"github.com/gin-gonic/gin"
)

type UserView struct {
	userService *services.UserService
}

func NewUserView(userService *services.UserService) *UserView {
	return &UserView{
		userService: userService,
	}
}

func (v *UserView) GetMeHandler(c *gin.Context) {
	uid, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

	user, err := v.userService.GetUserByID(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
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
