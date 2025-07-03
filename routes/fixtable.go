package routes

import (
	"foodieMatch/controller"
	"foodieMatch/firebase"

	"github.com/gin-gonic/gin"
)

func FixTableRoutes(r *gin.Engine) {
	userRoutes := r.Group("/data")
	userRoutes.Use(firebase.AuthMiddleware()) 
	userRoutes.GET("/getCusines", controller.GetCusines)	
	userRoutes.GET("/getDietyPrefrence", controller.GetDietyPrefrence)
	userRoutes.GET("/getFoodAllergy", controller.GetFoodAllergy)
	userRoutes.GET("/getDishes", controller.GetDishes)
	userRoutes.GET("/getCookingStyle", controller.GetCookingStyle)
	userRoutes.GET("/getSpiceLevel", controller.GetSpiceLevel)
	userRoutes.GET("/getGalleryTag", controller.GetGalleryTag)
	userRoutes.GET("/getCookingLevel", controller.GetCookingLevel)
	userRoutes.GET("/getPrompt", controller.GetPrompt)
	userRoutes.GET("/getDiningStyle", controller.GetDiningStyle)
	userRoutes.GET("/getFoodHabit", controller.GetFoodHabit)
	userRoutes.GET("/getFavDrink", controller.GetFavDrink)
}