package routes


import (
	"foodieMatch/controller"
	"foodieMatch/firebase"

	"github.com/gin-gonic/gin"
)

func FoodPrefrenceRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user/foodprefrence")
	userRoutes.Use(firebase.AuthMiddleware()) 
	userRoutes.POST("/createFoodPreference", controller.CreateFoodPreference)
	userRoutes.GET("/getFoodPreference", controller.GetFoodPreference)
	userRoutes.PUT("/updateFoodPreference", controller.UpdateFoodPreference)
	userRoutes.PUT("/updateFavoriteCuisine", controller.UpdateFavoriteCuisine)
	userRoutes.PUT("/updateDietaryPreference", controller.UpdateDietaryPreference)
	userRoutes.PUT("/updateFoodAllergy", controller.UpdateFoodAllergy)
	userRoutes.PUT("/updateFavoriteDish", controller.UpdateFavoriteDish)
	userRoutes.PUT("/updateCookingStyle", controller.UpdateCookingStyle)
	userRoutes.PUT("/updateSpiceTolerance", controller.UpdateSpiceTolerance)
}