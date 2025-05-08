package routes

import (
	"foodieMatch/controller"
	"foodieMatch/firebase"

	"github.com/gin-gonic/gin"
)

func UserPostRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user/post")
	userRoutes.Use(firebase.AuthMiddleware()) 
	userRoutes.POST("/createUserPost", controller.CreateUserPost)
	userRoutes.GET("/getAllUserPost", controller.GetAllUserPost)
	userRoutes.GET("/getUserPostById/:post_id", controller.GetUserPostById)
	userRoutes.PUT("/updateUserPost", controller.UpdateUserPost)
	userRoutes.PUT("/updateUserTag", controller.UpdateUserTag)
	userRoutes.DELETE("/deleteUserPost/:post_id", controller.DeleteUserPost)
}