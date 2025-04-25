package controller

import (
	"fmt"
	"foodieMatch/db"
	"foodieMatch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
func CreateFoodPreference(c *gin.Context) {

	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var foodPreference models.FoodPrefrenceRequestDTO
	if err := c.ShouldBindJSON(&foodPreference); err != nil {
		fmt.Println("Error binding food preference data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	userID := uid.(string)

	err := db.InsertUserFoodPreference(userID, &foodPreference)
	fmt.Println("Food preference: ", foodPreference)
	if err != nil {
		fmt.Println("Error inserting food preference: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert food preference"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food preferences created successfully"})	

}

func GetFoodPreference(c *gin.Context) {
	uid, exists := c.Get("uid")	
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	foodPreference, err := db.GetUserFoodPreference(userID)

	if err != nil && err.Error() == "not found" {
		fmt.Println("No food preferences found for user")
		c.JSON(http.StatusNotFound, gin.H{"error": "Food preferences not found"})
		return
	}
	if err != nil {
		fmt.Println("Error getting food preference: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get food preference"})
		return
	}	

	c.JSON(http.StatusOK, foodPreference)
}
	
	
func UpdateFoodPreference(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var updateFoodPreference models.UserFoodPreference
	if err := c.ShouldBindJSON(&updateFoodPreference); err != nil {
		fmt.Println("Error binding food preference data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := db.UpdateUserFoodPreference(userID, &updateFoodPreference)
	if err != nil {
		fmt.Println("Error updating food preference: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food preference"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food preferences updated successfully"})
}

func UpdateFavoriteCuisine(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}		

	userID := uid.(string)

	var updateFavoriteCuisine models.FavoriteCuisineRequestDTO
	if err := c.ShouldBindJSON(&updateFavoriteCuisine); err != nil {
		fmt.Println("Error binding favorite cuisine data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateUserFavoriteCuisine(userID, &updateFavoriteCuisine)
	if err != nil {
		fmt.Println("Error updating favorite cuisine: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update favorite cuisine"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite cuisine updated successfully"})
}

func UpdateDietaryPreference(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var updateDietaryPreference models.DietaryPreferenceRequestDTO
	if err := c.ShouldBindJSON(&updateDietaryPreference); err != nil {
		fmt.Println("Error binding dietary preference data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateUserDietaryPreference(userID, &updateDietaryPreference)
	if err != nil {
		fmt.Println("Error updating dietary preference: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update dietary preference"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dietary preference updated successfully"})
}

func UpdateFoodAllergy(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var updateFoodAllergy models.FoodAllergyRequestDTO
	if err := c.ShouldBindJSON(&updateFoodAllergy); err != nil {
		fmt.Println("Error binding food allergy data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	err := db.UpdateUserFoodAllergy(userID, &updateFoodAllergy)
	if err != nil {
		fmt.Println("Error updating food allergy: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food allergy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food allergy updated successfully"})
}

func UpdateFavoriteDish(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var updateFavoriteDish models.FavoriteDishRequestDTO
	if err := c.ShouldBindJSON(&updateFavoriteDish); err != nil {
		fmt.Println("Error binding favorite dish data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateUserFavoriteDish(userID, &updateFavoriteDish)
	if err != nil {
		fmt.Println("Error updating favorite dish: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update favorite dish"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite dish updated successfully"})	
}

func UpdateCookingStyle(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var updateCookingStyle models.CookingStyleRequestDTO
	if err := c.ShouldBindJSON(&updateCookingStyle); err != nil {
		fmt.Println("Error binding cooking style data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateUserCookingStyle(userID, &updateCookingStyle)
	if err != nil {
		fmt.Println("Error updating cooking style: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cooking style"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cooking style updated successfully"})
}

func UpdateSpiceTolerance(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var updateSpiceTolerance models.SpiceToleranceRequestDTO
	if err := c.ShouldBindJSON(&updateSpiceTolerance); err != nil {
		fmt.Println("Error binding spice tolerance data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateUserSpiceTolerance(userID, &updateSpiceTolerance)
	if err != nil {
		fmt.Println("Error updating spice tolerance: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update spice tolerance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Spice tolerance updated successfully"})	
}