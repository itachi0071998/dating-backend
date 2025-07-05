package s3

import (
	"context"
	"log"
	"os"

	// "os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func init() {
	log.Println("Connection to the S3 Server started!")
	customCreds := aws.NewCredentialsCache(
		credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"), config.WithCredentialsProvider(customCreds))
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}
	S3Client = s3.NewFromConfig(cfg)
	log.Println("Connected to S3 Server!")
}
