package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"log"
)

var DB *dynamo.DB

func InitDb() *dynamo.DB {
	awsConfig := &aws.Config{
		Region:      aws.String("us-east-1"), // Especifique a região
		Credentials: credentials.NewStaticCredentials("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", ""),
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	if sess == nil {
		log.Fatal("AWS session is nil")
	}

	db := dynamo.New(sess)

	if db == nil {
		log.Fatal("DynamoDB client is nil")
	}

	return db
}
