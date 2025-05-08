package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"foodieMatch/db"
	"foodieMatch/models"
)


func CreateUserPost(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	var userPostRequestDTO models.UserPostRequestDTO
	if err := c.ShouldBindJSON(&userPostRequestDTO); err != nil {
		fmt.Println("Error binding user post data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	postID, err := db.CreateUserPost(userID, userPostRequestDTO)
	if err != nil {
		fmt.Println("Error creating user post: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User post created successfully", "post_id": postID})
}

func GetAllUserPost(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	userPosts, err := db.GetAllUserPost(userID)
	if err != nil {
		fmt.Println("Error getting user posts: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User posts retrieved successfully", "posts": userPosts})
}

func GetUserPostById(c *gin.Context) {
	postID := c.Param("post_id")

	fmt.Println("Post ID: ", postID)

	userPost, err := db.GetUserPostById(postID)

	if err != nil && err.Error()=="post not found" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err != nil {
		fmt.Println("Error getting user post: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User post retrieved successfully", "post": userPost})
}
	

func UpdateUserPost(c *gin.Context) {
	var userPostRequestDTO models.UpdateUserPostRequestDTO
	if err := c.ShouldBindJSON(&userPostRequestDTO); err != nil {
		fmt.Println("Error binding user post data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	err := db.UpdateUserPost(userPostRequestDTO)
	if err != nil {
		fmt.Println("Error updating user post: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User post updated successfully"})
}

func DeleteUserPost(c *gin.Context) {
	postID := c.Param("post_id")

	err := db.DeleteUserPost(postID)

	if err != nil {
		fmt.Println("Error deleting user post: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User post deleted successfully"})

}

func UpdateUserTag(c *gin.Context) {
	var updatePostTag models.UpdatePostTagRequestDTO	
	if err := c.ShouldBindJSON(&updatePostTag); err != nil {
		fmt.Println("Error binding user tag data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateUserPostTag(&updatePostTag)
	if err != nil {
		fmt.Println("Error updating user tag: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User tag updated successfully"})
}
