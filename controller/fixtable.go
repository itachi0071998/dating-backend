package controller

import (
	"foodieMatch/db"
	"github.com/gin-gonic/gin"
	"net/http"
)


func GetCusines(c *gin.Context) {
	cuisines, err := db.GetCusines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cuisines"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cuisines retrieved successfully", "cuisines": cuisines})
}	

func GetDietyPrefrence(c *gin.Context) {
	dietaryPreferences, err := db.GetDietyPrefrence()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dietary preferences"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dietary preferences retrieved successfully", "dietaryPreferences": dietaryPreferences})	
}

func GetFoodAllergy(c *gin.Context) {
	foodAllergies, err := db.GetFoodAllergy()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get food allergies"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Food allergies retrieved successfully", "foodAllergies": foodAllergies})	
}

func GetDishes(c *gin.Context) {
	dishes, err := db.GetDishes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dishes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dishes retrieved successfully", "dishes": dishes})	
}	

func GetCookingStyle(c *gin.Context) {
	cookingStyles, err := db.GetCookingStyle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cooking styles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cooking styles retrieved successfully", "cookingStyles": cookingStyles})	
}		

func GetSpiceLevel(c *gin.Context) {
	spiceLevels, err := db.GetSpiceLevel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get spice levels"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Spice levels retrieved successfully", "spiceLevels": spiceLevels})	
}		

func GetGalleryTag(c *gin.Context) {
	galleryTags, err := db.GetGalleryTag()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gallery tags"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gallery tags retrieved successfully", "galleryTags": galleryTags})	
}	

func GetCookingLevel(c *gin.Context) {
	cookingLevels, err := db.GetCookingLevel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cooking levels"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cooking levels retrieved successfully", "cookingLevels": cookingLevels})	
}		

func GetPrompt(c *gin.Context) {
	prompts, err := db.GetPrompt()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prompts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prompts retrieved successfully", "prompts": prompts})	
}

func GetDiningStyle(c *gin.Context) {
	diningStyles, err := db.GetDiningStyle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dining styles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dining styles retrieved successfully", "diningStyles": diningStyles})	
}

func GetFoodHabit(c *gin.Context) {
	foodHabits, err := db.GetFoodHabit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get food habits"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Food habits retrieved successfully", "foodHabits": foodHabits})	
}

func GetFavDrink(c *gin.Context) {
	favDrinks, err := db.GetFavDrink()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fav drinks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fav drinks retrieved successfully", "favDrinks": favDrinks})	
}