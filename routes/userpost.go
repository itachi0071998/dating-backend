package routes

import (
	"foodieMatch/controller"
	"foodieMatch/firebase"

	"github.com/gin-gonic/gin"
)

func UserPostRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	userRoutes.Use(firebase.AuthMiddleware()) 
	userRoutes.POST("/post", controller.CreateUserPost)
	userRoutes.GET("/post", controller.GetAllUserPost)
	userRoutes.GET("/post/:post_id", controller.GetUserPostById)
	userRoutes.PUT("/post/:post_id", controller.UpdateUserPost)
	userRoutes.DELETE("/post/:post_id", controller.DeleteUserPost)
}