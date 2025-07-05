package controller

import (
	"fmt"
	"foodieMatch/models"
	"foodieMatch/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserProfile(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var userProfile models.UserProfileRequest
	if err := c.ShouldBindJSON(&userProfile); err != nil {
		fmt.Println("Error binding user profile data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uid.(string)

	exists, err := service.UserExists(userID)
	if err != nil {
		fmt.Println("Error checking user existence: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
		return
	}
	if exists {
		fmt.Println("User already exists")
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	log.Println("User does not exist")

	err = service.InsertUserProfile(userID, &userProfile)
	fmt.Println("User profile: ", userProfile)
	if err != nil {
		fmt.Println("Error inserting user profile: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User profile created successfully"})

}

func UpdateUserProfile(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var userProfile models.UserProfileRequest
	if err := c.ShouldBindJSON(&userProfile); err != nil {
		fmt.Println("Error binding user profile data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uid.(string)

	exists, err := service.UserExists(userID)
	if err != nil {
		fmt.Println("Error checking user existence: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
		return
	}
	if !exists {
		fmt.Println("User does not exist")
		c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
		return
	}
	err = service.UpdateUserProfile(userID, &userProfile)
	fmt.Println("User profile: ", userProfile)
	if err != nil {
		fmt.Println("Error updating user profile: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})	
}

func GetUserProfile(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	userProfile, err := service.GetUserProfile(userID)
	if err != nil {
		fmt.Println("Error getting user profile: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userProfile": userProfile})
}