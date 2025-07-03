package routes

import (
	"foodieMatch/controller"
	"foodieMatch/firebase"
	"github.com/gin-gonic/gin"
)

func UserPromptRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	userRoutes.Use(firebase.AuthMiddleware()) 
	userRoutes.POST("/prompt", controller.CreateUserPrompt)
	userRoutes.GET("/prompt", controller.GetUserPrompt)
	userRoutes.PUT("/prompt", controller.UpdateUserPrompt)
}