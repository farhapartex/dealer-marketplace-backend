package user

import (
	"github.com/farhapartex/dealer-marketplace-be/apps/user/middleware"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/services"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/views"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	authService := services.NewAuthService()
	authView := views.NewAuthView(authService)

	userService := services.NewUserService()
	userView := views.NewUserView(userService)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authView.SignupHandler)
		authGroup.POST("/verify", authView.VerifyCodeHandler)
		authGroup.POST("/signin", authView.SigninHandler)
	}

	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/me", userView.GetMeHandler)
	}
}
