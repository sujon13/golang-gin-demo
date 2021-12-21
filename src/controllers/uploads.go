package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func Upload(c *gin.Context) {
	session := c.MustGet("session").(*session.Session)
	uploader := s3manager.NewUploader(session)
	bucket := GetEnvWithKey("BUCKET_NAME")
	region := GetEnvWithKey("AWS_REGION")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		//ACL:    aws.String("public-read"),
		Key:  aws.String(header.Filename),
		Body: file,
	})

	if err != nil {
		log.Println("Error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload file",
		})
		return
	}

	// https://bucket-name.s3.Region.amazonaws.com/key-name
	filepath := "https://" + bucket + ".s3." + region + ".amazonaws.com/" + header.Filename
	c.JSON(
		http.StatusOK,
		gin.H{
			"filepath": filepath,
		},
	)
}

func MultipleFileUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	if len(files) == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "No files found for key: files",
			},
		)
		return
	}

	fileNames := make([]string, len(files))
	for ind, file := range files {
		log.Println(ind, file.Filename)
		destination, _ := os.Create("../static/" + file.Filename)

		err := c.SaveUploadedFile(file, destination.Name())
		if err != nil {
			log.Println("File: " + destination.Name() + " could not be saved for")
			fileNames[ind] = ""
		} else {
			fileNames[ind] = destination.Name()
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": fileNames})
}
