package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"log"
)

func InitDb() *dynamo.DB {
	awsConfig := &aws.Config{
		Region: aws.String("us-east-1"),
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
