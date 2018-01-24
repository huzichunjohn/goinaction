package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Printf("s3Bucket: %s, secretKey: %s\n", s3Bucket, secretKey)
}