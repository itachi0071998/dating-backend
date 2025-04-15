package routes

import (
	"foodieMatch/controller"
	"foodieMatch/firebase"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user/auth")
	userRoutes.Use(firebase.AuthMiddleware()) 
	userRoutes.GET("/getUser", controller.GetUserProfile)
	userRoutes.POST("/createUser", controller.CreateUser)
	userRoutes.PUT("/updateUser", controller.UpdateUser)
}