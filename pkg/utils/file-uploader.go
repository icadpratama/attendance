package utils

import (
	"github.com/minio/minio-go/v6"
	log "github.com/icadpratama/attendance/internal/logger"
	"github.com/icadpratama/attendance/pkg/utils"
)

func upload() {
	endpoint := utils.MustGet("MINIO_ENDPOINT")
	accessKeyID := utils.MustGet("MINIO_ACCESS_KEY")
	secretAccessKey := utils.MustGet("MINIO_SECRET_KEY")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Info(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Info("We already own %s\n", bucketName)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Info("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	objectName := "golden-oldies.zip"
	filePath := "/tmp/golden-oldies.zip"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType:contentType})
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Successfully uploaded %s of size %d\n", objectName, n)
}