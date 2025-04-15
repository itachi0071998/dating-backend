package controller

import (
	"fmt"
	"foodieMatch/db"
	"foodieMatch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Get UID from context set by auth middleware
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Error binding user data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the user ID from the authenticated token
	user.ID = uid.(string)

	// Validate user data
	if err := user.Validate(); err != nil {
		fmt.Println("Error validating user data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	exists, err := db.CheckUserExists(user.ID)
	if err != nil {
		fmt.Println("Error checking user existence: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
		return
	}
	if exists {
		fmt.Println("User already exists: ", user.ID)
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	fmt.Println("User: ", user)

	// Save the new user
	if err := db.SaveUser(&user); err != nil {
		fmt.Println("Error creating user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func UpdateUser(c *gin.Context) {
	// Get UID from context set by auth middleware
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		fmt.Println("Error binding user data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the user ID from the authenticated token
	updatedUser.ID = uid.(string)

	// Check if user exists
	exists, err := db.CheckUserExists(updatedUser.ID)
	if err != nil {
		fmt.Println("Error checking user existence: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
		return
	}
	if !exists {
		fmt.Println("User not found: ", updatedUser.ID)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Validate updated user data
	if err := updatedUser.Validate(); err != nil {
		fmt.Println("Error validating user data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user
	if err := db.UpdateUser(&updatedUser); err != nil {
		fmt.Println("Error updating user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}
