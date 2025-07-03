package controller

import 	(
	"fmt"
	"net/http"
	"foodieMatch/models"
	"foodieMatch/db"
	"github.com/gin-gonic/gin"
)

func CreateUserPrompt(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var createUserPrompt []models.PromptDTO
	if err := c.ShouldBindJSON(&createUserPrompt); err != nil {
		fmt.Println("Error binding user prompts data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.InsertUserPrompt(userID, createUserPrompt)
	if err != nil {
		fmt.Println("Error creating user prompts: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user prompts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User prompts created successfully"})
}

func UpdateUserPrompt(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var updateUserPrompt models.PromptUpdateRequestDTO
	if err := c.ShouldBindJSON(&updateUserPrompt); err != nil {
		fmt.Println("Error binding user prompts data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateUserPrompt(userID, updateUserPrompt)
	if err != nil {
		fmt.Println("Error updating user prompts: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user prompts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User prompts updated successfully"})
}

func GetUserPrompt(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	prompt, err := db.GetUserPrompt(userID)
	if err != nil {
		fmt.Println("Error getting user prompts: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user prompts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User prompts retrieved successfully", "data": prompt})
}