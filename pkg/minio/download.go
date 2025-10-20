package minio

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
)

func DownloadObject(ObjName, BucketName string, client *minio.Client) error {
	object, err := client.GetObject(context.Background(), BucketName,
		ObjName, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("DownloadObject error: %s", err)
	}
	defer object.Close()

	filePath := fmt.Sprintf("/tmp/%s", ObjName)
	localFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("CreateFile error: %s", err)
	}
	defer localFile.Close()

	if _, err = io.Copy(localFile, object); err != nil {
		return fmt.Errorf("ioCopy error: %s", err)
	}
	fmt.Printf("File downloaded in: %s", filePath)
	return nil
}
