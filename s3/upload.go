package s3

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"path/filepath"
// 	"time"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
// )

// func UploadHandler(w http.ResponseWriter, r *http.Request) {
// 	// 10 MB max file
// 	r.ParseMultipartForm(10 << 20)

// 	file, handler, err := r.FormFile("file")
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error reading file: %v", err), http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	// give it a unique name in S3
// 	key := fmt.Sprintf("uploads/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))

// 	uploader := manager.NewUploader(S3Client)
// 	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
// 		Bucket: aws.String("foodie-match"),
// 		Key:    aws.String(key),
// 		Body:   file,
// 	})		
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error uploading: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// construct S3 URL
// 	s3URL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", "foodie-match", "us-east-1", key)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(fmt.Sprintf(`{"url":"%s"}`, s3URL)))
// }
