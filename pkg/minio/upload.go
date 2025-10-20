package minio

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func UploadObject(ObjBytes []byte, ObjName, BucketName string, client *minio.Client) error {

	uploadInfo, err := client.PutObject(context.Background(), BucketName,
		ObjName,
		bytes.NewReader(ObjBytes),
		int64(len(ObjBytes)), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return fmt.Errorf("UploadObject error %s:", err)
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
	return nil
}
