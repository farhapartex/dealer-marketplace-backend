package views

import (
	"net/http"

	"github.com/farhapartex/dealer-marketplace-be/apps/user/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/services"
	"github.com/farhapartex/dealer-marketplace-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthView struct {
	authService *services.AuthService
}

func NewAuthView(authService *services.AuthService) *AuthView {
	return &AuthView{
		authService: authService,
	}
}

func (v *AuthView) SignupHandler(c *gin.Context) {
	var payload dtos.CreateUserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"message": err.Error(),
		})
		return
	}

	response, err := v.authService.CreateNewUser(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully. Please verify your account.",
		"data":    response,
	})
}

func (v *AuthView) VerifyCodeHandler(c *gin.Context) {
	var payload dtos.VerifyCodeRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"message": err.Error(),
		})
		return
	}

	userID, err := utils.VerifyJWT(payload.Token)
	if err != nil || userID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Verification failed",
			"message": "Invalid code",
		})
		return
	}

	if err := v.authService.VerifyAuthCode(userID, payload.Code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Verification failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User verified successfully. Your account is now active.",
	})
}

func (v *AuthView) SigninHandler(c *gin.Context) {
	var payload dtos.SigninRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"message": err.Error(),
		})
		return
	}

	response, err := v.authService.Signin(payload)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Authentication failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
