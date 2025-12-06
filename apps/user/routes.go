package user

import (
	"github.com/farhapartex/dealer-marketplace-be/apps/user/services"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/views"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	authService := services.NewAuthService()
	authView := views.NewAuthView(authService)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authView.SignupHandler)
		authGroup.POST("/verify", authView.VerifyCodeHandler)
		authGroup.POST("/signin", authView.SigninHandler)
	}
}
