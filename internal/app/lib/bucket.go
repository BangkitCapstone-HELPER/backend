package lib

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

type ClientUploader struct {
	Cl         *storage.Client
	projectID  string
	bucketName string
}

func NewBucket(cfg config.BucketConfig) ClientUploader {
	fmt.Println("Starting client connection")
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", cfg.Path())
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return ClientUploader{
		Cl:         client,
		bucketName: cfg.BucketName(),
		projectID:  cfg.ProjectID(),
	}
}

func (c *ClientUploader) UploadFile(file multipart.File, folder string, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.Cl.Bucket(c.bucketName).Object(folder + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
