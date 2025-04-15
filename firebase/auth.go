package firebase

import (
	"context"
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		
		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		if idToken == authHeader{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is missing"})
			c.Abort()
			return
		}

		client, err := App.Auth(context.Background())
		 if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Initializing firebase auth: "+err.Error()})
			c.Abort()
			return
		 }
		 token, err := client.VerifyIDToken(context.Background(), idToken)
		 
		 if err != nil {
			log.Println("invalid token", err, token)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid IdToken"})
			c.Abort()
			return
		 }
			
		// Store both the token and UID in the context
		c.Set("userToken", token)
		c.Set("uid", token.UID)

		// Continue to the next handler
		c.Next()
	}
}
