package routes

import (
	"foodieMatch/controller"
	"foodieMatch/firebase"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	userRoutes.Use(firebase.AuthMiddleware()) 
	userRoutes.POST("/profile", controller.CreateUserProfile)
	userRoutes.PUT("/profile", controller.UpdateUserProfile)
	userRoutes.GET("/profile", controller.GetUserProfile)
	userRoutes.POST("/profile/upload/:bucket", controller.UploadUserProfileHandler)
}