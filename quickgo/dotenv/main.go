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

  // now do something with s3 or whatever
  fmt.Println(s3Bucket, secretKey)
}