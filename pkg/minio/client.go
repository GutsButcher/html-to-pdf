package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinIOClient(ep, accessKey, secretKey string, SSL bool) (*minio.Client, error) {
	endpoint := ep
	accessKeyID := accessKey
	secretAccessKey := secretKey
	useSSL := SSL

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("minIO client init error: %s", err)
	}

	fmt.Printf("%#v\n", minioClient) // minioClient is now set up

	return minioClient, nil
}
