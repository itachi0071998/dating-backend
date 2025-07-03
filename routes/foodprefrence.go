package routes


import (
	// "foodieMatch/controller"
	"foodieMatch/firebase"

	"github.com/gin-gonic/gin"
)

func FoodPrefrenceRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	userRoutes.Use(firebase.AuthMiddleware()) 
	// userRoutes.POST("/foodPrefrence", controller.CreateFoodPreference)
	// userRoutes.GET("/foodPrefrence", controller.GetFoodPreference)
	// userRoutes.PUT("/foodPrefrence", controller.UpdateFoodPreference)
	// userRoutes.PUT("/foodPrefrence/updateFavoriteCuisine", controller.UpdateFavoriteCuisine)
	// userRoutes.PUT("/foodPrefrence/updateDietaryPreference", controller.UpdateDietaryPreference)
	// userRoutes.PUT("/foodPrefrence/updateFoodAllergy", controller.UpdateFoodAllergy)
	// userRoutes.PUT("/foodPrefrence/updateFavoriteDish", controller.UpdateFavoriteDish)
	// userRoutes.PUT("/foodPrefrence/updateCookingStyle", controller.UpdateCookingStyle)
	// userRoutes.PUT("/foodPrefrence/updateSpiceTolerance", controller.UpdateSpiceTolerance)
	// userRoutes.PUT("/foodPrefrence/updateFavDrink", controller.UpdateFavDrink)
	// userRoutes.PUT("/foodPrefrence/updateCookingLevel", controller.UpdateCookingLevel)
	// userRoutes.PUT("/foodPrefrence/updateDiningStyles", controller.UpdateDiningStyles)
	// userRoutes.PUT("/foodPrefrence/updateFoodHabits", controller.UpdateFoodHabits)
}