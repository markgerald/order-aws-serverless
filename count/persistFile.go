package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"time"
)

func PersistFile(itemCount int) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your AWS region
	})
	uploader := s3manager.NewUploader(sess)

	csvData := new(bytes.Buffer)
	writer := csv.NewWriter(csvData)
	if err := writer.Write([]string{"ItemCount", fmt.Sprintf("%d", itemCount)}); err != nil {
		log.Fatalf("Error writing to CSV: %v", err)
	}
	writer.Flush()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(time.DateTime + "-orders-count.csv"),
		Body:   bytes.NewReader(csvData.Bytes()),
	})
	if err != nil {
		log.Fatalf("Unable to upload %q to %q, %v", "item-count.csv", "your-bucket-name", err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", "item-count.csv", "your-bucket-name")
}
