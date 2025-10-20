package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func CreateBucket(bucketName string, client *minio.Client) error {
	// Create a bucket at region 'us-east-1' with object locking enabled.
	err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("CreateBucket error: %s", err)
	}
	fmt.Printf("Successfully created %s\n", bucketName)
	return nil
}
