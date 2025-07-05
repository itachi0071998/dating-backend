package controller

import (
	"context"
	"fmt"
	"foodieMatch/s3"
	"net/http"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

const maxFileSize = 50 << 20 // 50 MB

func UploadUserProfileHandler(c *gin.Context) {

	uid, exists := c.Get("uid")
	if !exists {
		fmt.Println("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := uid.(string)

	bucket := c.Param("bucket")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("error reading file: %v", err)})
		return
	}

	// file size check
	if file.Size > maxFileSize {
		c.JSON(400, gin.H{"error": "file exceeds 50MB limit"})
		return
	}

	// open file
	f, err := file.Open()
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("error opening file: %v", err)})
		return
	}
	defer f.Close()

	// s3 key
	key := fmt.Sprintf("%s/%s/%d%s", userID, bucket, time.Now().UnixNano(), filepath.Ext(file.Filename))

	uploader := manager.NewUploader(s3.S3Client)

	_, err = uploader.Upload(context.TODO(), &awss3.PutObjectInput{
		Bucket: aws.String("foodie-user-profile"),
		Key:    aws.String(key),
		Body:   f,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("error uploading to s3: %v", err)})
		return
	}

	// s3 url
	s3Region := "us-east-1"
	s3URL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", "foodie-user-profile", s3Region, key)

	c.JSON(200, gin.H{
		"url": s3URL,
	})

}
